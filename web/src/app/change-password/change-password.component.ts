import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { User } from '../user/user.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  user = new User(0, {});
  newPassword = '';
  repeatPassword = '';
  oldPassword = '';

  constructor(private title: Title, private router: Router) { }

  ngOnInit() {
    Config.setLogin(true);
    Config.API('user', { username: localStorage.getItem('username') }).subscribe(values => this.initUser(values));
  }

  initUser(values: any) {
    this.user = new User(values['id'], values);
    this.title.setTitle(Config.lang('changePassword') + ' - ' + Config.get('title'));
  }

  changePassword() {
    if (this.oldPassword.trim() === '' || this.newPassword.trim() === '' || this.repeatPassword.trim() === '') {
      Config.openSnackBar(Config.lang('fillAllFields'));
    } else if (this.newPassword !== this.repeatPassword) {
      Config.openSnackBar(Config.lang('passwordsNotMatch'));
    } else if (Config.hash(this.oldPassword) !== localStorage.getItem('password')) {
      Config.openSnackBar(Config.lang('wrongPassword'));
    } else if (this.newPassword.length < 8) {
      Config.openSnackBar(Config.lang('passwordMinLength'));
    } else {
      Config.API('editpassword', {
        username: localStorage.getItem('username'), password: localStorage.getItem('password'), newPassword: Config.hash(this.newPassword)
      }).subscribe(values => values['success'] === true ? this.passwordChanged() :
        values['error'] === '403' ? Config.openSnackBar(Config.lang('wrongPassword')) : Config.openSnackBar('error'));
    }
  }

  passwordChanged() {
    localStorage.setItem('password', Config.hash(this.newPassword));
    Config.openSnackBar(Config.lang('passwordChanged'));
    this.router.navigate(['/edit-profile']);
  }
}
