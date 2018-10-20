import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';
import { Language } from '../language';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  constructor() { }

  ngOnInit() {
  }
}
