import { sha3_512 } from 'js-sha3';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material';

export namespace Config {
  let triedLogin = false;
  export let openSnackBar;
  export let login: boolean;
  export let router: Router;
  export let captcha: string;
  export let http: HttpClient;
  export let snackBar: MatSnackBar;
  export const baseUrl = environment.production
    ? '/'
    : 'http://localhost:1813/';
  export const apiUrl = baseUrl + 'api/';
  const configMap: Map<string, string> = new Map<string, string>();
  const languageMap: Map<string, string> = new Map<string, string>();

  export function get(key: string): string {
    return configMap.get(key);
  }

  export function lang(name: string): string {
    return languageMap.get(name);
  }

  export function loadLanguage(title: Title, site: string) {
    let language = localStorage.getItem('language');
    if (language === null) {
      language = get('language');
    }
    API('lang', {}).subscribe(values => Object.entries(values[language]).forEach(([key, value]) => setLang(key, value as string, title, site)));
    if (site === null) {
      title.setTitle(get('title'));
    }
  }

  function setLang(key: string, value: string, title: Title, site: string) {
    languageMap.set(key, value);
    if (key === site) {
      title.setTitle(value + ' - ' + get('title'));
    }
  }

  export function getCaptcha() {
    API('newcaptcha', {}).subscribe(values => captcha = values['captcha']);
  }

  export function loadFirst(title: Title, site: string) {
    API('conf', {}).subscribe(values => Object.entries(values).forEach(([key, value]) => loadFirstSet(key, value as string, title, site)));
  }

  let langOrTitleSet = false;
  function loadFirstSet(key: string, value: string, title: Title, site: string) {
    configMap.set(key, value);
    if (key === 'title' || key === 'language') {
      if (langOrTitleSet) {
        loadLanguage(title, site);
      }
      langOrTitleSet = true;
    }
  }

  export function API(url: string, body: any): Observable<any> {
    return http.post<any>(apiUrl + url, body);
  }

  export function setLogin(title: Title, site: string, redirect: boolean) {
    if (!triedLogin) {
      loadFirst(title, site);
      API('login', { username: localStorage.getItem('username'), password: localStorage.getItem('password') })
        .subscribe(values => validateLogin(values, redirect));
    } else {
      title.setTitle(lang(site) + ' - ' + get('title'));
      if (redirect && !login) {
        router.navigate(['/']);
        openSnackBar(lang('loginRequired'));
      }
    }
  }

  function validateLogin(values: any, redirect: boolean) {
    login = values['valid'];
    if (redirect && !login) {
      router.navigate(['/']);
    }
    triedLogin = true;
  }

  export function getUsername(): string {
    return localStorage.getItem('username');
  }

  export function getPassword(): string {
    return localStorage.getItem('password');
  }

  export function logout() {
    localStorage.removeItem('username');
    localStorage.removeItem('password');
    login = false;
  }

  export function registeredDate(registered: Object): string {
    return createdDate((new Date(<string>registered)).getTime() / 1000);
  }

  export function createdDate(created: Object): string {
    const date = new Date(<number>created * 1000);
    return (date.getDate() <= 9 ? '0' + date.getDate() : date.getDate()) + '.' +
      ((date.getMonth() + 1) <= 9 ? '0' + (date.getMonth() + 1) : (date.getMonth() + 1)) + '.' + date.getFullYear();
  }

  export function time(created: Object): string {
    const date = new Date(<number>created * 1000);
    return (date.getHours() <= 9 ? '0' + date.getHours() : date.getHours()) + ':' +
      (date.getMinutes() <= 9 ? '0' + date.getMinutes() : date.getMinutes());
  }

  export function hash(text: string): string {
    return sha3_512('gorum_' + sha3_512(text));
  }
}
