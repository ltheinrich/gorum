import { Component, OnInit } from '@angular/core';
import { User } from '../user/user.component';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  users: User[] = [];

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(this.title, 'users', false, null);
    Config.API('users', {}).subscribe(values => this.listUsers(values));
  }

  listUsers(values: any) {
    Object.entries(values).forEach(user => this.users.push(new User(<number>(<unknown>user[0]), <{ [key: string]: string }>user[1])));
  }
}
