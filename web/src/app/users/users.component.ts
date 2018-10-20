import { Component, OnInit } from '@angular/core';
import { User } from '../user/user.component';
import { Config } from '../config';
import { Language } from '../language';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  conf = Config;
  lang = Language;

  users: User[] = [];

  constructor(private title: Title) { }

  ngOnInit() {
    this.title.setTitle(Language.get('users') + ' - ' + Config.get('title'));
    Config.API('users', {}).subscribe(values =>
      Object.entries(values).forEach(user =>
        this.users.push(new User(<number><unknown>user[0], <{ [key: string]: Object; }>user[1]))));
  }

  registeredDate(registered: Object): string {
    const date = new Date(<string>registered);
    return date.getDate() + '.' + (date.getMonth() + 1) + '.' + date.getFullYear();
  }

}
