import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-page-not-found',
  templateUrl: './page-not-found.component.html',
  styleUrls: ['./page-not-found.component.css']
})
export class PageNotFoundComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;
  constructor(private title: Title) { }
  ngOnInit() {
    Config.setLogin(this.title, 'pageDoesNotExist', false, null);
  }
}
