import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';
import { ActivatedRoute } from '@angular/router';
import { Thread } from '../thread/thread.component';

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
  lang = Config.lang;

  user = new User(0, {});
  id = +this.route.snapshot.paramMap.get('id');
  threads: Thread[] = [];

  constructor(private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    Config.API('user', { userID: this.id }).subscribe(values => this.initUser(values));
    Config.API('lastuserthreads', { userID: this.id })
      .subscribe(values => this.listThreads(values));
  }

  listThreads(values: any) {
    Object.entries(values).forEach(thread => this.threads.push(new Thread(<number>thread[1]['id'], <string>thread[1]['name'],
      <string>thread[1]['board'], /* <number>thread[1]['author'] */ null, <number>thread[1]['created'], /* <string>thread[1]['content'] */
      null, <string>thread[1]['authorName'], <string>thread[1]['authorAvatar'], <number>thread[1]['answer'])));
    this.threads.sort((a, b) => b.answer - a.answer);
  }

  initUser(values: any) {
    this.user = new User(this.id, values);
    this.title.setTitle(this.user.data['username'] + ' - ' + Config.get('title'));
  }
}
