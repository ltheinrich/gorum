import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { Language } from '../language';
import { User } from '../user/user.component';
import { appInstance } from '../app.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  user = new User(0, {});
  newPassword = '';
  repeatPassword = '';
  oldPassword = '';

  constructor(
    private title: Title,
    private router: Router
  ) { }

  ngOnInit() {
    Config.setLogin(true);
    this.title.setTitle(Language.get('changePassword') + ' - ' + Config.get('title'));
    Config.API('user', { username: localStorage.getItem('username') })
      .subscribe(values => this.initUser(values));
  }

  initUser(values: any) {
    this.user = new User(values['id'], values);
    this.title.setTitle(
      this.user.data['username'] + ' - ' + Config.get('title')
    );
  }

  changePassword() {
    if (this.oldPassword.trim() === '' || this.newPassword.trim() === '' || this.repeatPassword.trim() === '') {
      appInstance.openSnackBar(Language.get('fillAllFields'));
    } else if (this.newPassword !== this.repeatPassword) {
      appInstance.openSnackBar(Language.get('passwordsNotMatch'));
    } else if (Config.hash(this.oldPassword) !== localStorage.getItem('password')) {
      appInstance.openSnackBar(Language.get('wrongPassword'));
    } else {
      Config.API('editpassword', {
        username: localStorage.getItem('username'),
        password: localStorage.getItem('password'),
        newPassword: Config.hash(this.newPassword)
      }).subscribe(values => values['success'] === true ? this.passwordChanged() :
        values['error'] === '403' ? appInstance.openSnackBar(Language.get('wrongPassword')) : appInstance.openSnackBar('error'));
    }
  }

  passwordChanged() {
    localStorage.setItem('password', Config.hash(this.newPassword));
    appInstance.openSnackBar(Language.get('passwordChanged'));
    this.router.navigate(['/edit-profile']);
  }

}
