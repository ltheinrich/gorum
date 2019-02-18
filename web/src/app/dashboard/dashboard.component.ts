import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';
import { Thread } from '../thread/thread.component';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  threads: Thread[] = [];
  lastThreadsShown = true;

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(this.title, 'dashboard', false, null);
    Config.API('lastthreads', { username: Config.getUsername() }).subscribe(values => this.listThreads(values));
  }

  listThreads(values: any) {
    Object.entries(values).forEach(thread => this.threads.push(new Thread(<number>thread[1]['id'],
      <string>thread[1]['name'], <number>thread[1]['board'], null, <number>thread[1]['author'], <number>thread[1]['created'],
      /* <string>thread[1]['content'] */ null, <string>thread[1]['authorName'], <string>thread[1]['authorAvatar'],
      <number>thread[1]['answer'])));
    this.threads.sort((a, b) => b.answer - a.answer);
  }
}
