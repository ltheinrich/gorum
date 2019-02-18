import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { Config } from '../config';

export class Thread {
  id: number;
  name: string;
  board: number;
  boardName: string;
  author: number;
  created: number;
  content: string;
  authorName: string;
  authorAvatar: string;
  answer: number;

  constructor(id: number, name: string, board: number, boardName: string, author: number, created: number,
    content: string, authorName: string, authorAvatar: string, answer: number) {
    this.id = id;
    this.name = name;
    this.board = board;
    this.boardName = boardName;
    this.author = author;
    this.created = created;
    this.content = content;
    this.authorName = authorName;
    this.authorAvatar = authorAvatar;
    this.answer = answer;
  }
}

export class Post {
  id: number;
  author: number;
  authorName: string;
  authorAvatar: string;
  content: string;
  created: number;

  constructor(id: number, author: number, authorName: string, authorAvatar: string, content: string, created: number) {
    this.id = id;
    this.author = author;
    this.authorName = authorName;
    this.authorAvatar = authorAvatar;
    this.content = content;
    this.created = created;
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
  lang = Config.lang;

  thread = new Thread(0, null, null, null, null, null, null, null, null, null);
  posts: Post[] = [];
  id = +this.route.snapshot.paramMap.get('id');
  captcha: string;

  constructor(private router: Router, private route: ActivatedRoute, private title: Title) { }

  ngOnInit() {
    Config.API('thread', { threadID: this.id }).subscribe(values => this.initThread(values));
    Config.API('posts', { threadID: this.id }).subscribe(values => this.listPosts(values));
    Config.getCaptcha();
  }

  initThread(values: any) {
    this.thread = new Thread(<number>values['id'], <string>values['name'], <number>values['board'], <string>values['boardName'],
      <number>values['author'], <number>values['created'], <string>values['content'], <string>values['authorName'],
      <string>values['authorAvatar'], null);
    if (this.thread.name !== undefined) {
      Config.setLogin(this.title, 'thread', false, this.thread.name);
    } else {
      Config.setLogin(this.title, 'thread', false, null);
    }
  }

  listPosts(values: any) {
    this.posts = [];
    Object.entries(values).forEach(post => this.listPost(post));
    this.posts.sort((a, b) => a.created - b.created);
  }

  listPost(post: any) {
    if (<number>post[1]['id'] >= 1) {
      this.posts.push(new Post(<number>post[1]['id'], <number>post[1]['author'],
        <string>post[1]['authorName'], <string>post[1]['authorAvatar'], <string>post[1]['content'], <number>post[1]['created']));
    }
  }

  deleteThread() {
    Config.API('deletethread', { username: Config.getUsername(), token: Config.getToken(), threadID: this.id })
      .subscribe(values => this.afterDeleteThread(values));
  }

  deletePost(id: number) {
    Config.API('deletepost', { username: Config.getUsername(), token: Config.getToken(), postID: id })
      .subscribe(values => this.afterDeletePost(values));
  }

  afterDeletePost(values: any) {
    if (values['done'] === true) {
      Config.API('posts', { threadID: this.id }).subscribe(valuesPosts => this.listPosts(valuesPosts));
      Config.openSnackBar(Config.lang('postDeleted'));
    } else if (values['error'] !== undefined) {
      Config.openSnackBar(values['error']);
    }
  }

  afterDeleteThread(values: any) {
    if (values['done'] === true) {
      Config.openSnackBar(Config.lang('threadDeleted'));
      this.router.navigate(['/board/' + this.thread.board]);
    } else if (values['error'] !== undefined) {
      Config.openSnackBar(values['error']);
    }
  }

  post(content: string) {
    Config.API('newpost', {
      username: Config.getUsername(), token: Config.getToken(), thread: this.id, content: content,
      captcha: Config.captcha, captchaValue: this.captcha
    }).subscribe(values => this.proccessResponse(values));
  }

  proccessResponse(values: any) {
    if (values['error'] === '400') {
      Config.openSnackBar(Config.lang('fillAllFields'));
    } else if (values['error'] === '403') {
      Config.openSnackBar(Config.lang('wrongLogin'));
    } else if (values['error'] === '403 captcha') {
      Config.openSnackBar(Config.lang('wrongCaptcha'));
      Config.getCaptcha();
      this.captcha = '';
    } else if (values['error'] === '411') {
      Config.openSnackBar(Config.lang('contentMaxLength'));
    } else if (values['error'] !== undefined) {
      Config.openSnackBar(values['error']);
      Config.getCaptcha();
      this.captcha = '';
    } else {
      Config.openSnackBar(Config.lang('postCreated'));
      Config.getCaptcha();
      this.captcha = '';
      const element = <any>document.querySelector('trix-editor');
      element.editor.setSelectedRange([0, (<string>element.editor.getDocument().toString()).length]);
      element.editor.deleteInDirection('forward');
      Config.API('posts', { threadID: this.id }).subscribe(posts => this.listPosts(posts));
    }
  }
}
