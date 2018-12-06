import { Component, OnInit, Input } from '@angular/core';
import { Language } from '../language';
import { Config } from '../config';
import { ActivatedRoute } from '@angular/router';
import { Board } from '../board/board.component';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-new-thread',
  templateUrl: './new-thread.component.html',
  styleUrls: ['./new-thread.component.css']
})
export class NewThreadComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Language.get;

  id = +this.route.snapshot.paramMap.get('id');
  board: Board;
  threadTitle: string;

  constructor(private route: ActivatedRoute,
    private title: Title) { }

  ngOnInit() {
    Config.setLogin(true);
    Config.API('board', { boardID: this.id })
      .subscribe(values => this.initBoard(values));
  }

  initBoard(values: any) {
    this.board = new Board(values['id'], values['name'], values['description'], values['icon'], values['sort']);
    this.title.setTitle(
      Language.get('newThread') + ' - ' + Config.get('title')
    );
  }

  publish(content: string) {

  }
}
