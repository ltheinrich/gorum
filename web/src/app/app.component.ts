import { Component, ChangeDetectorRef, OnDestroy, Inject, OnInit } from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { Config } from './config';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog, MatSnackBar } from '@angular/material';
import { HttpClient } from '@angular/common/http';
import { Language } from './language';

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

  constructor(private http: HttpClient, public dialog: MatDialog, public snackBar: MatSnackBar,
    changeDetectorRef: ChangeDetectorRef, media: MediaMatcher) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this.mobileQueryListener);
  }

  ngOnInit(): void {
    Config.setHttp(this.http);
    Config.setLogin();
    Language.loadLanguage('de');
    Config.load(['title']);
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
    const dialogRef = this.dialog.open(LoginDialogOverview, { width: '300px', data: {} });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }

      if (result.username === undefined || result.password === undefined) {
        this.openSnackBar(Language.get('fillAllFields'));
      } else {
        const hashed = Config.hash(result.password);
        Config.API('login',
          { username: result.username, password: hashed })
          .subscribe(values => values['valid'] === true ?
            this.setLogin(result.username, hashed, Language.get('loginSuccess')) :
            this.openSnackBar(Language.get('loginWrong')));
      }
    });
  }

  openRegister(): void {
    const dialogRef = this.dialog.open(RegisterDialogOverview, { width: '300px', data: {} });
    dialogRef.afterClosed().subscribe(result => {
      if (result === undefined) {
        return;
      }

      if (result.username === undefined || result.mail === undefined || result.password === undefined || result.repeat === undefined) {
        this.openSnackBar(Language.get('fillAllFields'));
      } else if (result.password === result.repeat) {
        const hashed = Config.hash(result.password);
        Config.API('register', { username: result.username, mail: result.mail, password: hashed })
          .subscribe(values => values['done'] === true ?
            this.setLogin(result.username, hashed, Language.get('userCreated')) :
            this.openSnackBar(Language.get('userAlreadyExists')));
      } else {
        this.openSnackBar(Language.get('passwordsNotMatch'));
      }
    });
  }

  doLogout(): void {
    Config.logout();
    this.openSnackBar(Language.get('loggedOut'));
  }

  openSnackBar(message: string) {
    this.snackBar.open(message, Language.get('close'), { duration: 5000 });
  }
}

export interface LoginDialogData {
  username: string; password: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'login-dialog-overview', templateUrl: './login-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class LoginDialogOverview {
  conf = Config.get;
  lang = Language.get;
  constructor(public dialogRef: MatDialogRef<LoginDialogOverview>, @Inject(MAT_DIALOG_DATA) public data: LoginDialogData) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
}

export interface RegisterDialogData {
  username: string; mail: string; password: string; repeat: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'register-dialog-overview', templateUrl: './register-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class RegisterDialogOverview {
  conf = Config.get;
  lang = Language.get;
  constructor(public dialogRef: MatDialogRef<RegisterDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: RegisterDialogData) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
}
