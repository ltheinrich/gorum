import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Language } from '../language';
import { Title } from '@angular/platform-browser';
import { Board } from '../board/board.component';
import { iteratorToArray } from '@angular/animations/browser/src/util';

@Component({
  selector: 'app-boards',
  templateUrl: './boards.component.html',
  styleUrls: ['./boards.component.css']
})
export class BoardsComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  categoryNames: string[] = [];
  categories: Map<string, Board[]> = new Map<string, Board[]>();
  titleSet = false;

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    Config.API('boards', {}).subscribe(values =>
      Object.entries(values).forEach(category =>
        Object.entries(category[1]).forEach(board => this.addBoard(<string><unknown>category[0],
          new Board(board[1]['id'], board[1]['name'], board[1]['description'], board[1]['icon']))))
    );
  }

  addBoard(category: string, board: Board) {
    if (!this.titleSet) {
      this.title.setTitle(Language.get('boards') + ' - ' + Config.get('title'));
      this.titleSet = true;
    }

    if (!this.categories.has(category)) {
      this.categoryNames.push(category);
      this.categories.set(category, []);
    }
    this.categories.get(category).push(board);
  }

}
