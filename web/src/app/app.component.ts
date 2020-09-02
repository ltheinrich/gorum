import { Component, ChangeDetectorRef, OnDestroy, Inject, OnInit } from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { Config } from './config';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpClient } from '@angular/common/http';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';

export let appInstance: AppComponent;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  private mobileQueryListener: () => void;
  mobileQuery: MediaQueryList;
  footer: string;

  constructor(private http: HttpClient, private title: Title, public dialog: MatDialog, public snackBar: MatSnackBar,
    changeDetectorRef: ChangeDetectorRef, media: MediaMatcher, private router: Router) {
    appInstance = this;
    Config.snackBar = this.snackBar;
    Config.openSnackBar = this.openSnackBar;
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this.mobileQueryListener);

  }

  ngOnInit() {
    Config.http = this.http;
    Config.router = this.router;
    Config.API('footer', {}).subscribe(values => this.setFooter(values));
  }

  ngOnDestroy() {
    this.mobileQuery.removeListener(this.mobileQueryListener);
  }

  private setFooter(values: any) {
    if (values['footer'] !== undefined) {
      this.footer = values['footer'];
    }
  }

  private setLogin(username: string, token: string, message: string) {
    localStorage.setItem('username', username);
    localStorage.setItem('token', token);
    Config.login = true;
    Config.openSnackBar(message);
  }

  changeLanguage(language: string) {
    localStorage.setItem('language', language);
    Config.loadLanguage(this.title, undefined, null);
  }

  openLogin() {
    Config.getCaptcha();
    const dialogRef = this.dialog.open(LoginDialogOverview, { width: '300px', data: {} });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }
      this.login(result, dialogRef, null);
    });
  }

  public login(result: any, dialogRef: MatDialogRef<any>, data: any) {
    if (result.username === undefined || result.password === undefined || (result.captcha === undefined && Config.captcha !== undefined)) {
      Config.openSnackBar(Config.lang('fillAllFields'));
      return;
    } else if (result.username.length > 32) {
      Config.openSnackBar(Config.lang('usernameMaxLength'));
      return;
    } else if (result.password.length < 8) {
      Config.openSnackBar(Config.lang('passwordMinLength'));
      return;
    } else {
      const hashed = Config.hash(result.password);
      Config.API('login', { username: result.username, password: hashed, captcha: Config.captcha, captchaValue: data.captcha })
        .subscribe(values => this.closeDialogOnLogin(values, result.username, hashed, dialogRef, data));
    }
  }

  openRegister() {
    Config.getCaptcha();
    const dialogRef = this.dialog.open(RegisterDialogOverview, { width: '300px', data: {} });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }
      this.register(result, dialogRef, null);
    });
  }

  public register(result: any, dialogRef: MatDialogRef<any>, data: any) {
    if (
      result.username === undefined || result.password === undefined || result.repeat === undefined ||
      (result.captcha === undefined && Config.captcha !== undefined)) {
      Config.openSnackBar(Config.lang('fillAllFields'));
      return;
    } else if (result.username.length > 32) {
      Config.openSnackBar(Config.lang('usernameMaxLength'));
      return;
    } else if (result.password.length < 8) {
      Config.openSnackBar(Config.lang('passwordMinLength'));
      return;
    } else if (result.password === result.repeat) {
      const hashed = Config.hash(result.password);
      Config.API('register', {
        username: result.username, password: hashed, captcha: Config.captcha, captchaValue: result.captcha
      }).subscribe(values => this.closeDialogOnLogin(values, result.username, hashed, dialogRef, data));
    } else {
      Config.openSnackBar(Config.lang('passwordsNotMatch'));
    }
  }

  closeDialogOnLogin(values: any, username: string, hashed: string, dialogRef: MatDialogRef<any>, data: any) {
    if (values['token'] !== undefined && values['done'] !== true) {
      this.setLogin(username, values['token'], Config.lang('loginSuccess'));
      dialogRef.close();
    } else if (values['error'] === '403') {
      Config.openSnackBar(Config.lang('loginWrong'));
    } else if (values['done'] === true) {
      this.setLogin(username, values['token'], Config.lang('userCreated'));
      dialogRef.close();
    } else if (values['error'] === '400') {
      Config.openSnackBar(Config.lang('wrongData'));
    } else if (values['error'] === '403 captcha') {
      Config.openSnackBar(Config.lang('wrongCaptcha'));
    } else {
      Config.openSnackBar(Config.lang('userAlreadyExists'));
    }
    Config.getCaptcha();
    data.captcha = '';
  }

  doLogout() {
    Config.logout();
    Config.openSnackBar(Config.lang('loggedOut'));
  }

  openSnackBar(message: string) {
    Config.snackBar.open(message, Config.lang('close'), { duration: 4000 });
  }
}

export interface LoginDialogData {
  username: string; password: string; captcha: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'login-dialog-overview',
  templateUrl: './login-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class LoginDialogOverview {
  config = Config;
  conf = Config.get;
  lang = Config.lang;
  constructor(public dialogRef: MatDialogRef<LoginDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: LoginDialogData) { }
  onNoClick() {
    this.dialogRef.close();
  }
  doLogin() {
    appInstance.login(this.data, this.dialogRef, this.data);
  }
}

export interface RegisterDialogData {
  username: string; password: string; repeat: string; captcha: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'register-dialog-overview',
  templateUrl: './register-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class RegisterDialogOverview {
  config = Config;
  conf = Config.get;
  lang = Config.lang;
  constructor(public dialogRef: MatDialogRef<RegisterDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: RegisterDialogData) { }
  onNoClick() {
    this.dialogRef.close();
  }
  doRegister() {
    appInstance.register(this.data, this.dialogRef, this.data);
  }
}
