import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { ActivatedRoute, Router } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-edit-thread',
  templateUrl: './edit-thread.component.html',
  styleUrls: ['./edit-thread.component.css']
})
export class EditThreadComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  id = +this.route.snapshot.paramMap.get('id');
  threadTitle: string;
  threadContent: string;

  constructor(private route: ActivatedRoute, private title: Title, private router: Router) { }

  ngOnInit() {
    Config.API('thread', { threadID: this.id }).subscribe(values => this.initThread(values));
  }

  initThread(values: any) {
    if (values['authorName'] !== Config.getUsername() || values['created'] === undefined) {
      this.router.navigate(['/thread/' + this.id]);
    }
    this.threadTitle = values['name'];
    this.threadContent = values['content'];
    const element = <any>document.querySelector('trix-editor');
    element.editor.insertHTML(this.threadContent);
    Config.setLogin(this.title, 'editThread', true, this.threadTitle);
  }

  publish(content: string) {
    if (this.threadTitle.length > 32) {
      Config.openSnackBar(Config.lang('titleMaxLength'));
    } else {
      Config.API('editthread', {
        username: Config.getUsername(), token: Config.getToken(), threadID: this.id, title: this.threadTitle, content: content,
      }).subscribe(values => this.proccessResponse(values));
    }
  }

  proccessResponse(values: any) {
    if (values['error'] === '400') {
      Config.openSnackBar(Config.lang('fillAllFields'));
    } else if (values['error'] === '403') {
      Config.openSnackBar(Config.lang('wrongLogin'));
    } else if (values['error'] === '411') {
      Config.openSnackBar(Config.lang('contentMaxLength'));
    } else if (values['error'] !== undefined) {
      Config.openSnackBar(values['error']);
    } else {
      Config.openSnackBar(Config.lang('threadEdited'));
      this.router.navigate(['/thread/' + this.id]);
    }
  }
}
