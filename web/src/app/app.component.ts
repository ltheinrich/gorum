import { Component, ChangeDetectorRef, OnDestroy, Inject } from '@angular/core';
import { MediaMatcher } from '@angular/cdk/layout';
import { Config } from './config';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog, MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnDestroy {
  conf = Config.Get;

  username: string;
  password: string;
  repeat: string;
  mobileQuery: MediaQueryList;

  private mobileQueryListener: () => void;

  constructor(public dialog: MatDialog,
    changeDetectorRef: ChangeDetectorRef, media: MediaMatcher,
    public snackBar: MatSnackBar) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this.mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this.mobileQueryListener);
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this.mobileQueryListener);
  }

  openLogin(): void {
    const dialogRef = this.dialog.open(LoginDialogOverview, {
      width: '300px',
      data: { username: this.username, password: this.password }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result.username === undefined && result.password === undefined) {
        this.openSnackBar('Bitte fülle alle Eingabefelder aus', 'Schließen');
        return;
      }

      this.username = result.username;
      this.password = result.password;
    });
  }

  openRegister(): void {
    const dialogRef = this.dialog.open(RegisterDialogOverview, {
      width: '300px',
      data: { username: this.username, password: this.password, repeat: this.repeat }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result.username === undefined && result.password === undefined && result.repeat === undefined) {
        this.openSnackBar('Bitte fülle alle Eingabefelder aus', 'Schließen');
        return;
      }

      this.username = result.username;
      this.password = result.password;
      this.repeat = result.repeat;
    });
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 4000,
    });
  }

}

export interface LoginDialogData {
  username: string;
  password: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'login-dialog-overview',
  templateUrl: './login-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class LoginDialogOverview {

  constructor(public dialogRef: MatDialogRef<LoginDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: LoginDialogData) { }

  onNoClick(): void {
    this.dialogRef.close();
  }

}

export interface RegisterDialogData {
  username: string;
  password: string;
  repeat: string;
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'register-dialog-overview',
  templateUrl: './register-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class RegisterDialogOverview {

  constructor(public dialogRef: MatDialogRef<RegisterDialogOverview>,
    @Inject(MAT_DIALOG_DATA) public data: RegisterDialogData) { }

  onNoClick(): void {
    this.dialogRef.close();
  }

}
