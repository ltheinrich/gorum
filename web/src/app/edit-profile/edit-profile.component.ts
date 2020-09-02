import { Component, OnInit } from '@angular/core';
import { User, UserData } from '../user/user.component';
import { Config } from '../config';
import { Title } from '@angular/platform-browser';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
  config = Config;
  conf = Config.get;
  lang = Config.lang;

  user = new User(0, {});
  userData = new UserData({});
  userDataOld = new UserData({});

  updateRequests = 0;
  finishedUpdateRequests = 0;
  errorMessage = '';

  constructor(private router: Router, private title: Title, public dialog: MatDialog) { }

  ngOnInit() {
    Config.setLogin(this.title, 'editProfile', true, null);
    Config.API('user', { username: Config.getUsername() }).subscribe(values => this.user = new User(values['id'], values));
    Config.API('userdata', {
      dataNames: ['website', 'eMailAddress', 'mastodon', 'twitter', 'youtube', 'wire', 'discord', 'aboutMe'],
      username: Config.getUsername()
    }).subscribe(values => this.initUserData(values));
  }

  initUserData(values: any) {
    this.userData.userData = values;
    this.userDataOld.userData = JSON.parse(JSON.stringify(values));
    const element = <any>document.querySelector('trix-editor');
    element.editor.insertHTML(<string>this.userData.userData['aboutMe']);
  }

  saveProfile(aboutMe: string) {
    if (aboutMe) {
      this.userData.userData['aboutMe'] = aboutMe;
    }

    const newUsername = <string>this.user.data['username'];
    if (Config.getUsername() !== newUsername) {
      if (newUsername === '') {
        Config.openSnackBar(Config.lang('emptyUsername'));
        return;
      } else if (newUsername.length > 32) {
        Config.openSnackBar(Config.lang('usernameMaxLength'));
        return;
      }
    }
    const newWebsite = <string>this.userData.userData['website'];
    if (newWebsite) {
      if (this.toValidLink(newWebsite, false) == null) {
        Config.openSnackBar(Config.lang('websiteInvalid'));
        return;
      } else {
        this.userData.userData['website'] = this.toValidLink(newWebsite, false);
      }
    }
    const newEmailAddress = <string>this.userData.userData['eMailAddress'];
    if (newEmailAddress) {
      if (!newEmailAddress.includes('@') || !newEmailAddress.includes('.') || newEmailAddress.length <= 5) {
        Config.openSnackBar(Config.lang('eMailAddressInvalid'));
        return;
      }
    }
    const newMastodon = <string>this.userData.userData['mastodon'];
    if (newMastodon) {
      if (this.toValidLink(newMastodon, false) == null) {
        Config.openSnackBar(Config.lang('mastodonLinkInvalid'));
        return;
      } else {
        this.userData.userData['mastodon'] = this.toValidLink(newMastodon, false);
      }
    }
    const newTwitter = <string>this.userData.userData['twitter'];
    if (newTwitter) {
      if (newTwitter.includes(' ') || newTwitter.length > 15) {
        Config.openSnackBar(Config.lang('twitterNameInvalid'));
        return;
      }
      if (newTwitter.includes('@')) {
        this.userData.userData['twitter'] = newTwitter.replace('@', '');
      }
    }
    const newYoutube = <string>this.userData.userData['youtube'];
    if (newYoutube) {
      if ((this.toValidLink(newYoutube, false) == null)) {
        Config.openSnackBar(Config.lang('youtubeLinkInvalid'));
        return;
      } else {
        this.userData.userData['youtube'] = this.toValidLink(newYoutube, true);
      }
    }
    const newWire = <string>this.userData.userData['wire'];
    if (newWire) {
      if (newWire.includes(' ')) {
        Config.openSnackBar(Config.lang('wireNameInvalid'));
        return;
      }
      if (newWire.includes('@')) {
        this.userData.userData['wire'] = newWire.replace('@', '');
      }
    }
    const newDiscord = <string>this.userData.userData['discord'];
    if (newDiscord) {
      if (newDiscord.includes(' ') || newDiscord.length < 6 || !newDiscord.includes('#')) {
        Config.openSnackBar(Config.lang('discordTagInvalid'));
        return;
      }
    }

    if (Config.getUsername() !== newUsername) {
      this.updateRequests++;
      Config.API('editusername', {
        username: Config.getUsername(), newUsername: newUsername, token: Config.getToken()
      }).subscribe(values => this.changedUsername(values, newUsername));
    }

    Object.keys(this.userData.userData).forEach(key => {
      if (this.userData.userData[key] !== this.userDataOld.userData[key]) {
        this.updateRequests++;
        Config.API('setuserdata', {
          dataName: key, dataValue: <string>this.userData.userData[key],
          username: Config.getUsername(), token: Config.getToken()
        }).subscribe(values => this.savedUserDataValue(values));
      }
    });

    if (this.updateRequests === 0) {
      this.savedUserDataValue({ 'success': true, 'status': 'nochanges' });
    }
  }

  savedUserDataValue(values: any) {
    this.finishedUpdateRequests++;
    if (values['success'] !== true) {
      const errorMessage = Config.lang(values['error']);
      this.errorMessage = errorMessage === undefined ? values['error'] : errorMessage;
    }
    if (values['status'] === 'nochanges') {
      this.router.navigate(['/user/' + this.user.id]);
      this.finishedUpdateRequests--;
      return;
    }
    if (this.updateRequests === this.finishedUpdateRequests && !this.errorMessage) {
      Config.openSnackBar(Config.lang('profileSaved'));
      this.router.navigate(['/user/' + this.user.id]);
    } else if (this.updateRequests === this.finishedUpdateRequests) {
      Config.openSnackBar(this.errorMessage);
    }
  }

  changedUsername(values: any, newUsername: string) {
    if (values['success'] === true) {
      localStorage.setItem('username', newUsername);
      Config.openSnackBar(Config.lang('changedUsername'));
      this.savedUserDataValue(values);
    } else {
      const errorMessage = Config.lang(values['error']);
      this.errorMessage = errorMessage === undefined ? values['error'] : errorMessage;
    }
    this.savedUserDataValue(values);
  }

  editAvatar(): void {
    this.dialog.open(AvatarDialogOverview, { width: '400px', data: {} });
  }

  toValidLink(link: string, https: boolean): string {
    if (!link.includes('.') || link.length <= 3) {
      return null;
    } else {
      if (!link.startsWith('http')) {
        if (https) {
          link = 'https://' + link;
        } else {
          link = 'http://' + link;
        }
      }
      if (link.startsWith('http://') && https) {
        link = link.replace('http://', 'https://');
      }
      if (link.includes(' ')) {
        link = link.replace(' ', '%20');
      }
    }
    return link;
  }
}

@Component({
  // tslint:disable-next-line:component-selector
  selector: 'avatar-dialog-overview',
  templateUrl: './avatar-dialog-overview.html'
})
// tslint:disable-next-line:component-class-suffix
export class AvatarDialogOverview {
  config = Config;
  username = btoa(Config.getUsername()).replace('+', '-').replace('/', '_').replace('=', '%3d');
  token = btoa(Config.getToken()).replace('+', '-').replace('/', '_').replace('=', '%3d');
  constructor(public dialogRef: MatDialogRef<AvatarDialogOverview>) { }
  onNoClick(): void {
    this.dialogRef.close();
  }
}
