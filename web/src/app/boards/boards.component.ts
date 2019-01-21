import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { Board } from '../board/board.component';

@Component({
  selector: 'app-boards',
  templateUrl: './boards.component.html',
  styleUrls: ['./boards.component.css']
})
export class BoardsComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  categoryNames: string[] = [];
  categories: Map<string, Board[]> = new Map<string, Board[]>();

  constructor(private title: Title) { }

  ngOnInit() {
    Config.setLogin(this.title, 'boards', false, null);
    Config.API('boards', {}).subscribe(values => this.addBoards(values));
  }

  addBoards(values: any) {
    Object.entries(values).forEach(category => Object.entries(category[1]).forEach(board => this.addBoard(<string><unknown>category[0],
      new Board(board[1]['id'], board[1]['name'], board[1]['description'], board[1]['icon'], board[1]['sort']))));
    this.categories.forEach(boards => boards.sort((a, b) => a.sort - b.sort));
    this.categoryNames.sort((a, b) => Array.from(this.categories.get(a))[0].sort - Array.from(this.categories.get(b))[0].sort);
  }

  addBoard(category: string, board: Board) {
    if (!this.categories.has(category)) {
      this.categoryNames.push(category);
      this.categories.set(category, []);
    }
    this.categories.get(category).push(board);
  }
}
