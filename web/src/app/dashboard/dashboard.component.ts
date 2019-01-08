import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    this.title.setTitle(Config.get('title'));
    Config.API('lastthreads', { username: Config.getUsername(), password: Config.getPassword() }).subscribe(values => console.log(values));
  }
}
