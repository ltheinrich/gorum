import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { AppComponent, RegisterDialogOverview } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatInputModule } from '@angular/material/input';
import { MatDialogModule } from '@angular/material/dialog';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatCardModule } from '@angular/material/card';
import { MatSelectModule } from '@angular/material/select';
import { LoginDialogOverview } from './app.component';
import { UsersComponent } from './users/users.component';
import { UserComponent } from './user/user.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { EditProfileComponent, AvatarDialogOverview } from './edit-profile/edit-profile.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { DeleteAccountComponent } from './delete-account/delete-account.component';
import { BoardsComponent } from './boards/boards.component';
import { ThreadComponent } from './thread/thread.component';
import { BoardComponent } from './board/board.component';
import { NewThreadComponent } from './new-thread/new-thread.component';
import { EditThreadComponent } from './edit-thread/edit-thread.component';
import { EditPostComponent } from './edit-post/edit-post.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { ExportDataComponent } from './export-data/export-data.component';
import { PageComponent } from './page/page.component';

@NgModule({
  declarations: [
    AppComponent, LoginDialogOverview, RegisterDialogOverview, AvatarDialogOverview, UsersComponent, UserComponent, DashboardComponent,
    EditProfileComponent, ChangePasswordComponent, BoardsComponent, ThreadComponent, BoardComponent, NewThreadComponent,
    EditThreadComponent, EditPostComponent, DeleteAccountComponent, PageNotFoundComponent, ExportDataComponent, PageComponent
  ],
  imports: [
    BrowserModule, HttpClientModule, BrowserAnimationsModule, FormsModule, AppRoutingModule, MatButtonModule, MatSelectModule,
    MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatInputModule, MatDialogModule, MatSnackBarModule, MatCardModule
  ],
  entryComponents: [LoginDialogOverview, RegisterDialogOverview, AvatarDialogOverview],
  exports: [],
  providers: [],
  bootstrap: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class AppModule { }
