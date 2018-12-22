import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { Language } from '../language';
import { Config } from '../config';

export class Thread {
  id: number;
  name: string;
  board: string;
  author: number;
  created: number;
  content: string;
  authorName: string;
  authorAvatar: string;

  constructor(id: number, name: string, board: string, author: number,
    created: number, content: string, authorName: string, authorAvatar: string) {
    this.id = id;
    this.name = name;
    this.board = board;
    this.author = author;
    this.created = created;
    this.content = content;
    this.authorName = authorName;
    this.authorAvatar = authorAvatar;
  }
}

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.css']
})
export class ThreadComponent implements OnInit {
  id = +this.route.snapshot.paramMap.get('id');

  constructor(private route: ActivatedRoute,
    private title: Title) { }

  ngOnInit() {
    this.title.setTitle(/* TODO: board title */ Language.get('board') + ' - ' + Config.get('title'));
  }

}
