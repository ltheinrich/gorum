import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { ActivatedRoute, Router } from '@angular/router';
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
  lang = Config.lang;

  id = +this.route.snapshot.paramMap.get('id');
  board: Board = new Board(null, null, null, null, null);
  threadTitle: string;
  captcha: string;

  constructor(private route: ActivatedRoute, private title: Title, private router: Router) { }

  ngOnInit() {
    Config.API('board', { boardID: this.id }).subscribe(values => this.initBoard(values));
    Config.getCaptcha();
  }

  initBoard(values: any) {
    if (values['name'] === undefined) {
      this.router.navigate(['/board/' + this.id]);
    }
    this.board = new Board(values['id'], values['name'], values['description'], values['icon'], values['sort']);
    Config.setLogin(this.title, 'newThread', true, values['name']);
  }

  publish(content: string) {
    if (this.threadTitle.length > 32) {
      Config.openSnackBar(Config.lang('titleMaxLength'));
    } else {
      Config.API('newthread', {
        username: Config.getUsername(), token: Config.getToken(), title: this.threadTitle, board: this.id,
        content: content, captcha: Config.captcha, captchaValue: this.captcha
      }).subscribe(values => this.proccessResponse(values));
    }
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
      Config.openSnackBar(Config.lang('threadCreated'));
      this.router.navigate(['/thread/' + values['id']]);
    }
  }
}
