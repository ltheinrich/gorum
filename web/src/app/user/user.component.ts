import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';
import { ActivatedRoute } from '@angular/router';
import { Thread } from '../thread/thread.component';

export class User {
  id: number;
  data: { [key: string]: string };

  constructor(id: number, data: { [key: string]: string }) {
    this.id = id;
    this.data = data;
  }
}

export class UserData {
  userData: { [key: string]: string };

  contactDataProvided: boolean;

  constructor(userData: { [key: string]: string }) {
    this.userData = userData;
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
  lang = Config.lang;

  user = new User(0, {});
  userData = new UserData(
    { 'website': '', 'eMailAddress': '', 'mastodon': '', 'twitter': '', 'youtube': '', 'wire': '', 'discord': '', 'aboutMe': '' });
  id = +this.route.snapshot.paramMap.get('id');
  threads: Thread[] = [];
  contactDataProvided: boolean;

  constructor(private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.API('user', { userID: this.id }).subscribe(values => this.initUser(values));
    Config.API('userdata',
      { dataNames: ['website', 'eMailAddress', 'mastodon', 'twitter', 'youtube', 'wire', 'discord', 'aboutMe'], userID: this.id })
      .subscribe(values => this.initUserData(values));
    Config.API('lastuserthreads', { userID: this.id })
      .subscribe(values => this.listThreads(values));
  }

  initUser(values: any) {
    this.user = new User(this.id, values);
    if (<string>values['username'] !== undefined) {
      Config.setLogin(this.title, 'user', false, <string>values['username']);
    } else {
      Config.setLogin(this.title, 'user', false, null);
    }
  }

  initUserData(values: any) {
    this.userData = new UserData(values);
    const contactData = ['website', 'eMailAddress', 'mastodon', 'twitter', 'youtube', 'wire', 'discord'];
    for (const entry of contactData) {
      if (this.userData.userData[entry]) {
        this.userData.contactDataProvided = true;
      }
    }
  }

  listThreads(values: any) {
    Object.entries(values).forEach(thread => this.threads.push(new Thread(<number>thread[1]['id'], <string>thread[1]['name'],
      <number>thread[1]['board'], null, /* <number>thread[1]['author'] */ null, <number>thread[1]['created'],
      /* <string>thread[1]['content'] */ null, <string>thread[1]['authorName'], <string>thread[1]['authorAvatar'],
      <number>thread[1]['answer'])));
    this.threads.sort((a, b) => b.answer - a.answer);
  }
}
