import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { Language } from '../language';
import { Config } from '../config';
import { appInstance } from '../app.component';

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
  config = Config;
  conf = Config.get;
  lang = Language.get;

  thread = new Thread(0, null, null, null, null, null, null, null);
  id = +this.route.snapshot.paramMap.get('id');

  constructor(private router: Router, private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.setLogin(false);
    Config.API('thread', { threadID: this.id }).subscribe(values => this.initThread(values));
  }

  initThread(values: any) {
    this.thread = new Thread(
      <number>values['id'], <string>values['name'], <string>values['board'], <number>values['author'],
      <number>values['created'], <string>values['content'], <string>values['authorName'], <string>values['authorAvatar']);
    this.title.setTitle(this.thread.name + ' - ' + Config.get('title'));
  }

  deleteThread() {
    Config.API('deletethread', { username: Config.getUsername(), password: Config.getPassword(), threadID: this.id })
      .subscribe(values => this.afterDeleteThread(values));
  }

  afterDeleteThread(values: any) {
    if (values['done'] === true) {
      appInstance.openSnackBar(Language.get('threadDeleted'));
      this.router.navigate(['/board/' + this.thread.board]);
    } else if (values['error'] !== undefined) {
      appInstance.openSnackBar(values['error']);
    }
  }
}
