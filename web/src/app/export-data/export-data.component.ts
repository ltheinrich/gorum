import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-export-data',
  templateUrl: './export-data.component.html',
  styleUrls: ['./export-data.component.css']
})
export class ExportDataComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(this.title, 'exportData', true, null);
    Config.API('exportdata', { username: Config.getUsername(), token: Config.getToken() }).subscribe(values => this.downloadData(values));
  }

  downloadData(values: any) {
    const element = document.getElementById('download');
    element.setAttribute('href', 'data:application/json;charset=utf-8,' + encodeURIComponent(JSON.stringify(values)));
    element.setAttribute('download', 'exported_data.json');
  }
}
