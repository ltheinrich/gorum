import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { AppComponent, RegisterDialogOverview } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import {
  MatButtonModule,
  MatToolbarModule,
  MatIconModule,
  MatSidenavModule,
  MatListModule,
  MatInputModule,
  MatDialogModule,
  MatSnackBarModule,
  MatCardModule
} from '@angular/material';
import { LoginDialogOverview } from './app.component';
import { UsersComponent } from './users/users.component';
import { UserComponent } from './user/user.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import {
  EditProfileComponent,
  AvatarDialogOverview
} from './edit-profile/edit-profile.component';
import { SafePipe } from './safepipe';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { BoardsComponent } from './boards/boards.component';
import { ThreadComponent } from './thread/thread.component';
import { BoardComponent } from './board/board.component';
import { FroalaEditorModule, FroalaViewModule } from 'angular-froala-wysiwyg';
import { NewThreadComponent } from './new-thread/new-thread.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginDialogOverview,
    RegisterDialogOverview,
    AvatarDialogOverview,
    UsersComponent,
    UserComponent,
    DashboardComponent,
    EditProfileComponent,
    SafePipe,
    ChangePasswordComponent,
    BoardsComponent,
    ThreadComponent,
    BoardComponent,
    NewThreadComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    AppRoutingModule,
    FroalaEditorModule.forRoot(),
    FroalaViewModule.forRoot(),
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    MatSidenavModule,
    MatListModule,
    MatInputModule,
    MatDialogModule,
    MatSnackBarModule,
    MatCardModule
  ],
  entryComponents: [
    LoginDialogOverview,
    RegisterDialogOverview,
    AvatarDialogOverview
  ],
  exports: [],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
