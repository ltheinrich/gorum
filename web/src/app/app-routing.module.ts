import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UsersComponent } from './users/users.component';
import { UserComponent } from './user/user.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { DeleteAccountComponent } from './delete-account/delete-account.component';
import { ThreadComponent } from './thread/thread.component';
import { BoardsComponent } from './boards/boards.component';
import { BoardComponent } from './board/board.component';
import { NewThreadComponent } from './new-thread/new-thread.component';
import { EditThreadComponent } from './edit-thread/edit-thread.component';
import { EditPostComponent } from './edit-post/edit-post.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { ExportDataComponent } from './export-data/export-data.component';
import { PageComponent } from './page/page.component';

export const routes: Routes = [
  { path: '', component: DashboardComponent, pathMatch: 'full' },
  { path: 'users', component: UsersComponent },
  { path: 'user/:id', component: UserComponent },
  { path: 'user/:id/:name', component: UserComponent },
  { path: 'edit-profile', component: EditProfileComponent },
  { path: 'change-password', component: ChangePasswordComponent },
  { path: 'delete-account', component: DeleteAccountComponent },
  { path: 'boards', component: BoardsComponent },
  { path: 'board/:id', component: BoardComponent },
  { path: 'board/:id/:name', component: BoardComponent },
  { path: 'thread/:id', component: ThreadComponent },
  { path: 'thread/:id/:name', component: ThreadComponent },
  { path: 'new-thread/:id', component: NewThreadComponent },
  { path: 'edit-thread/:id', component: EditThreadComponent },
  { path: 'edit-post/:id', component: EditPostComponent },
  { path: 'export-data', component: ExportDataComponent },
  { path: 'page/:name', component: PageComponent },
  { path: '**', component: PageNotFoundComponent }
];

@NgModule({ imports: [RouterModule.forRoot(routes)], exports: [RouterModule] })
export class AppRoutingModule { }
