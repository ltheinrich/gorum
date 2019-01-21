import { Component, OnInit } from '@angular/core';
import { Config } from '../config';
import { ActivatedRoute, Router } from '@angular/router';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-edit-post',
  templateUrl: './edit-post.component.html',
  styleUrls: ['./edit-post.component.css']
})
export class EditPostComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  id = +this.route.snapshot.paramMap.get('id');
  thread: number;
  postContent: string;

  constructor(private route: ActivatedRoute, private title: Title, private router: Router) { }

  ngOnInit() {
    Config.API('post', { postID: this.id }).subscribe(values => this.initPost(values));
  }

  initPost(values: any) {
    this.thread = values['thread'];
    this.postContent = values['content'];
    const element = <any>document.querySelector('trix-editor');
    element.editor.insertHTML(this.postContent);
    Config.setLogin(this.title, 'editPost', true, null);
  }

  publish(content: string) {
    Config.API('editpost', {
      username: Config.getUsername(), password: Config.getPassword(), postID: this.id, content: content,
    }).subscribe(values => this.proccessResponse(values));
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
      Config.openSnackBar(Config.lang('postEdited'));
      this.router.navigate(['/thread/' + this.thread]);
    }
  }
}
