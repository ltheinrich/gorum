import { Component, OnInit } from '@angular/core';
import { User } from '../user/user.component';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { MatDialogRef, MatDialog } from '@angular/material';
import { Router } from '@angular/router';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  user = new User(0, {});

  constructor(private router: Router, private title: Title, public dialog: MatDialog) { }

  ngOnInit() {
    Config.setLogin(this.title, 'editProfile', true, null);
    Config.API('user', { username: Config.getUsername() }).subscribe(values => this.initUser(values));
    Config.API('userdata', { dataNames: ['website'], username: Config.getUsername() })
      .subscribe(values => this.initUserData(values));
  }

  initUser(values: any) {
    this.user = new User(values['id'], values);
  }

  initUserData(values: any) {
    this.user.additionalData = values;
  }

  saveProfile() {
    // Change the userdata
    Config.API('setuserdata', {
      dataName: 'website', dataValue: <string>this.user.additionalData['website'],
      username: Config.getUsername(), token: Config.getToken()
    }).subscribe(values => this.savedProfile(values));

    // Change the username
    const newUsername = <string>this.user.data['username'];
    if (Config.getUsername() !== newUsername) {
      if (newUsername === '') {
        Config.openSnackBar(Config.lang('emptyUsername'));
      } else if (newUsername.length > 32) {
        Config.openSnackBar(Config.lang('usernameMaxLength'));
      } else {
        Config.API('editusername', {
          username: Config.getUsername(), newUsername: newUsername, token: Config.getToken()
        }).subscribe(values => this.changedUsername(values, newUsername));
      }
    } else {
      this.router.navigate(['/user/' + this.user.id]);
    }
  }

  savedProfile(values: any) {
    if (values['success'] === true) {
      Config.openSnackBar(Config.lang('profileSaved'));
      this.router.navigate(['/user/' + this.user.id]);
    } else {
      const errorMessage = Config.lang(values['error']);
      Config.openSnackBar(errorMessage === undefined ? values['error'] : errorMessage);
    }
  }

  changedUsername(values: any, newUsername: string) {
    if (values['success'] === true) {
      localStorage.setItem('username', newUsername);
      Config.openSnackBar(Config.lang('changedUsername'));
      this.router.navigate(['/user/' + this.user.id]);
    } else {
      const errorMessage = Config.lang(values['error']);
      Config.openSnackBar(errorMessage === undefined ? values['error'] : errorMessage);
    }
  }

  editAvatar(): void {
    this.dialog.open(AvatarDialogOverview, { width: '400px', data: {} });
  }
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'avatar-dialog-overview',
  templateUrl: './avatar-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class AvatarDialogOverview {
  config = Config;
  username = btoa(Config.getUsername());
  token = Config.getToken();
  constructor(public dialogRef: MatDialogRef<AvatarDialogOverview>) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
}
