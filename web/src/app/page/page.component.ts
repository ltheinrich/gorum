import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-page',
  templateUrl: './page.component.html',
  styleUrls: ['./page.component.css']
})
export class PageComponent implements OnInit {
  lang = Config.lang;
  page: string;
  name = this.route.snapshot.paramMap.get('name');

  constructor(private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.setLogin(this.title, 'dashboard', false, null);
    Config.API('page', { name: this.name }).subscribe(values => this.setPage(values));
  }

  private setPage(values: any) {
    if (values['page'] !== undefined) {
      this.page = values['page'];
    }
  }
}
