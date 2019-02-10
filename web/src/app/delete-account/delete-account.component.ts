import {Component, OnInit} from '@angular/core';
import {Config} from '../config';
import {Title} from '@angular/platform-browser';
import {User} from '../user/user.component';
import {Router} from '@angular/router';

@Component({
  selector: 'app-delete-account',
  templateUrl: './delete-account.component.html',
  styleUrls: ['./delete-account.component.css']
})
export class DeleteAccountComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  user = new User(0, {});
  password = '';

  constructor(private title: Title, private router: Router) {}

  ngOnInit() {
    Config.setLogin(this.title, 'deleteAccount', true, null);
  }

  deleteAccount() {
    if (this.password.trim() === '') {
      Config.openSnackBar(Config.lang('fillAllFields'));
    } else if (Config.hash(this.password) !== localStorage.getItem('password')) {
      Config.openSnackBar(Config.lang('wrongPassword'));
    } else {
      Config.API('deleteaccount', {
        username: localStorage.getItem('username'), password: localStorage.getItem('password')
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
