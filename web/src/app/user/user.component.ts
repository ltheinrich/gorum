import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Language } from '../language';
import { Config } from '../config';
import { ActivatedRoute } from '@angular/router';

export class User {
  id: number;
  data: { [key: string]: Object };

  constructor(id: number, data: { [key: string]: Object }) {
    this.id = id;
    this.data = data;
  }
}

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  user: User;
  id = +this.route.snapshot.paramMap.get('id');

  constructor(private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.API('user', { userID: this.id }).subscribe(values => this.initUser(values));
  }

  initUser(values: any) {
    this.user = new User(this.id, values);
    this.title.setTitle(this.user.data['username'] + ' - ' + Config.get('title'));
  }

}
