import { Component, OnInit } from '@angular/core';
import { User, UserData } from '../user/user.component';
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
  userData = new UserData({});
  userDataOld = new UserData({});

  updateRequests = 0;
  finishedUpdateRequests = 0;
  errorMessage = '';

  constructor(private router: Router, private title: Title, public dialog: MatDialog) { }

  ngOnInit() {
    Config.setLogin(this.title, 'editProfile', true, null);
    Config.API('user', { username: Config.getUsername() }).subscribe(values => this.user = new User(values['id'], values));
    Config.API('userdata', { dataNames: ['website', 'eMailAddress', 'twitter'], username: Config.getUsername() })
      .subscribe(values => this.initUserData(values));
  }

  initUserData(values: any) {
    this.userData.userData = values;
    this.userDataOld.userData = JSON.parse(JSON.stringify(values));
  }

  saveProfile() {
    const newUsername = <string>this.user.data['username'];
    if (Config.getUsername() !== newUsername) {
      if (newUsername === '') {
        Config.openSnackBar(Config.lang('emptyUsername'));
        return;
      } else if (newUsername.length > 32) {
        Config.openSnackBar(Config.lang('usernameMaxLength'));
        return;
      }
    }
    const newWebsite = <string>this.userData.userData['website'];
    if (newWebsite) {
      if (this.toValidLink(newWebsite) == null) {
        Config.openSnackBar(Config.lang('websiteInvalid'));
        return;
      } else {
        this.userData.userData['website'] = this.toValidLink(newWebsite);
      }
    }
    const newEmailAddress = <string>this.userData.userData['eMailAddress'];
    if (newEmailAddress) {
      if (!newEmailAddress.includes('@') || !newEmailAddress.includes('.') || newEmailAddress.length <= 5) {
        Config.openSnackBar(Config.lang('eMailAddressInvalid'));
        return;
      }
    }
    const newTwitter = <string>this.userData.userData['twitter'];
    if (newTwitter) {
      if (newTwitter.includes(' ') || newTwitter.length > 15) {
        Config.openSnackBar(Config.lang('twitterNameInvalid'));
        return;
      }
    }

    if (Config.getUsername() !== newUsername) {
      this.updateRequests++;
      Config.API('editusername', {
        username: Config.getUsername(), newUsername: newUsername, token: Config.getToken()
      }).subscribe(values => this.changedUsername(values, newUsername));
    }

    Object.keys(this.userData.userData).forEach(key => {
      if (this.userData.userData[key] !== this.userDataOld.userData[key]) {
        this.updateRequests++;
        Config.API('setuserdata', {
          dataName: key, dataValue: <string>this.userData.userData[key],
          username: Config.getUsername(), token: Config.getToken()
        }).subscribe(values => this.savedUserDataValue(values));
      }
    })

    if (this.updateRequests === 0) {
      this.savedUserDataValue({'success': true, 'status': 'nochanges'});
    }
  }

  savedUserDataValue(values: any) {
    this.finishedUpdateRequests++;
    if (values['success'] !== true) {
      const errorMessage = Config.lang(values['error']);
      this.errorMessage = errorMessage === undefined ? values['error'] : errorMessage;
    }
    if (values['status'] === 'nochanges') {
      this.router.navigate(['/user/' + this.user.id]);
      this.finishedUpdateRequests--;
      return;
    }
    if (this.updateRequests === this.finishedUpdateRequests && !this.errorMessage) {
      Config.openSnackBar(Config.lang('profileSaved'));
      this.router.navigate(['/user/' + this.user.id]);
    } else if (this.updateRequests === this.finishedUpdateRequests) {
      Config.openSnackBar(this.errorMessage);
    }
  }

  changedUsername(values: any, newUsername: string) {
    if (values['success'] === true) {
      localStorage.setItem('username', newUsername);
      Config.openSnackBar(Config.lang('changedUsername'));
      this.savedUserDataValue(values);
    } else {
      const errorMessage = Config.lang(values['error']);
      this.errorMessage = errorMessage === undefined ? values['error'] : errorMessage;
    }
    this.savedUserDataValue(values);
  }

  editAvatar(): void {
    this.dialog.open(AvatarDialogOverview, { width: '400px', data: {} });
  }

  toValidLink(link: string): string {
    if (!link.includes('.') || link.length <= 3) {
      return null;
    } else {
      if (!link.startsWith('http')) {
        link = 'http://' + link;
      }
      if (link.includes(' ')) {
        link = link.replace(' ', '%20');
      }
    }
    return link;
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
