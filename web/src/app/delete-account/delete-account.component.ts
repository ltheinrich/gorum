import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { User } from '../user/user.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-delete-account',
  templateUrl: './delete-account.component.html',
  styleUrls: ['./delete-account.component.css']
})
export class DeleteAccountComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  password = '';

  constructor(private title: Title, private router: Router) { }

  ngOnInit() {
    Config.setLogin(this.title, 'deleteAccount', true, null);
  }

  deleteAccount() {
    if (this.password.trim() === '') {
      Config.openSnackBar(Config.lang('fillAllFields'));
    } else {
      Config.API('deleteaccount', {
        username: Config.getUsername(), password: Config.hash(this.password), token: Config.getToken()
      }).subscribe(values => values['success'] === true ? this.accountDeleted() :
        values['error'] === '403' ? Config.openSnackBar(Config.lang('wrongPassword')) : Config.openSnackBar('error'));
    }
  }

  accountDeleted() {
    Config.logout();
    Config.openSnackBar(Config.lang('accountDeleted'));
    this.router.navigate(['/']);
  }
}
