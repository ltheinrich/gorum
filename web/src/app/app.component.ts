import {
  Component,
  ChangeDetectorRef,
  OnDestroy,
  Inject,
  OnInit
} from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { Config } from './config';
import {
  MatDialogRef,
  MAT_DIALOG_DATA,
  MatDialog,
  MatSnackBar
} from '@angular/material';
import { HttpClient } from '@angular/common/http';
import { Language } from './language';
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
  lang = Language.get;

  private mobileQueryListener: () => void;
  mobileQuery: MediaQueryList;

  constructor(
    private http: HttpClient,
    private title: Title,
    public dialog: MatDialog,
    public snackBar: MatSnackBar,
    changeDetectorRef: ChangeDetectorRef,
    media: MediaMatcher,
    private router: Router
  ) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this.mobileQueryListener);
    appInstance = this;
  }

  ngOnInit(): void {
    Config.setHttp(this.http);
    Config.loadFirst(['title'], this.title);
    Language.loadLanguage('de');
    Config.setRouter(this.router);
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this.mobileQueryListener);
  }

  private setLogin(username: string, password: string, message: string): void {
    localStorage.setItem('username', username);
    localStorage.setItem('password', password);
    Config.login = true;
    this.openSnackBar(message);
  }

  openLogin(): void {
    const dialogRef = this.dialog.open(LoginDialogOverview, {
      width: '300px',
      data: {}
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }
      this.login(result, dialogRef, null);
    });
  }

  public login(result: any, dialogRef: MatDialogRef<any>, data: any): void {
    if (result.username === undefined || result.password === undefined) {
      this.openSnackBar(Language.get('fillAllFields'));
      return;
    } else if (result.username.length > 32) {
      this.openSnackBar(Language.get('usernameMaxLength'));
      return;
    } else if (result.password.length < 8) {
      this.openSnackBar(Language.get('passwordMinLength'));
      return;
    } else {
      const hashed = Config.hash(result.password);
      Config.API('login', {
        username: result.username,
        password: hashed
      }).subscribe(values => this.closeDialogOnLogin(values, result.username, hashed, dialogRef, data));
    }
  }

  openRegister(): void {
    Config.getCaptcha();
    const dialogRef = this.dialog.open(RegisterDialogOverview, {
      width: '300px',
      data: {}
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }
      this.register(result, dialogRef, null);
    });
  }

  public register(result: any, dialogRef: MatDialogRef<any>, data: any): void {
    if (
      result.username === undefined || result.mail === undefined ||
      result.password === undefined || result.repeat === undefined ||
      (result.captcha === undefined && Config.captcha !== undefined)) {
      this.openSnackBar(Language.get('fillAllFields'));
      return;
    } else if (result.username.length > 32) {
      this.openSnackBar(Language.get('usernameMaxLength'));
      return;
    } else if (result.password.length < 8) {
      this.openSnackBar(Language.get('passwordMinLength'));
      return;
    } else if (result.password === result.repeat) {
      const hashed = Config.hash(result.password);
      Config.API('register', {
        username: result.username,
        mail: result.mail,
        password: hashed,
        captcha: Config.captcha,
        captchaValue: result.captcha
      }).subscribe(values => this.closeDialogOnLogin(values, result.username, hashed, dialogRef, data));
    } else {
      this.openSnackBar(Language.get('passwordsNotMatch'));
    }
  }

  closeDialogOnLogin(values: any, username: string, hashed: string, dialogRef: MatDialogRef<any>, data: any): void {
    if (values['valid'] === true) {
      this.setLogin(username, hashed, Language.get('loginSuccess'));
      dialogRef.close();
      return;
    } else if (values['valid'] === false) {
      this.openSnackBar(Language.get('loginWrong'));
    } else if (values['done'] === true) {
      this.setLogin(username, hashed, Language.get('userCreated'));
      dialogRef.close();
      return;
    } else if (values['error'] === '400') {
      this.openSnackBar(Language.get('wrongData'));
    } else if (values['error'] === '403 captcha') {
      this.openSnackBar(Language.get('wrongCaptcha'));
    } else {
      this.openSnackBar(Language.get('userAlreadyExists'));
    }
    Config.getCaptcha();
    data.captcha = '';
  }

  doLogout(): void {
    Config.logout();
    this.openSnackBar(Language.get('loggedOut'));
  }

  openSnackBar(message: string) {
    this.snackBar.open(message, Language.get('close'), { duration: 4000 });
  }
}

export interface LoginDialogData {
  username: string;
  password: string;
  captcha: string;
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
  lang = Language.get;
  constructor(
    public dialogRef: MatDialogRef<LoginDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: LoginDialogData
  ) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
  doLogin(): void {
    appInstance.login(this.data, this.dialogRef, this.data);
  }
}

export interface RegisterDialogData {
  username: string;
  mail: string;
  password: string;
  repeat: string;
  captcha: string;
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
  lang = Language.get;
  constructor(
    public dialogRef: MatDialogRef<RegisterDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: RegisterDialogData
  ) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
  doRegister(): void {
    appInstance.register(this.data, this.dialogRef, this.data);
  }
}
