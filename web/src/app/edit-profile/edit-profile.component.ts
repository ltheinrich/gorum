import { Component, OnInit } from '@angular/core';
import { User } from '../user/user.component';
import { Config } from '../config';
import { Language } from '../language';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { MatDialogRef, MatDialog } from '@angular/material';
import { PrivateSite } from '../private-site';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit, PrivateSite {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  user = new User(0, {});
  username = localStorage.getItem('username');

  constructor(
    private router: Router,
    private title: Title,
    public dialog: MatDialog
  ) { }

  ngOnInit() {
    Config.setLogin(true);
    Config.API('user', { username: localStorage.getItem('username') })
      .subscribe(values => this.initUser(values));
  }

  initUser(values: any) {
    this.user = new User(values['id'], values);
    this.title.setTitle(
      this.user.data['username'] + ' - ' + Config.get('title')
    );
  }

  saveProfile() { }

  editAvatar(): void {
    const dialogRef = this.dialog.open(AvatarDialogOverview, {
      width: '400px',
      data: {}
    });
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
  username = btoa(localStorage.getItem('username'));
  password = localStorage.getItem('password');
  constructor(public dialogRef: MatDialogRef<AvatarDialogOverview>) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
}
