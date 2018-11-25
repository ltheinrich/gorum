import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { Language } from '../language';
import { Config } from '../config';

export class Board {
  id: number;
  name: string;
  description: string;
  icon: string;

  constructor(id: number, name: string, description: string, icon: string) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.icon = icon;
  }
}

@Component({
  selector: 'app-board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {
  id = +this.route.snapshot.paramMap.get('id');

  constructor(private route: ActivatedRoute,
    private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    this.title.setTitle(/* TODO: board title */ Language.get('board') + ' - ' + Config.get('title'));
  }

}
