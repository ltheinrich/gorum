import { sha3_512 } from 'js-sha3';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';

export namespace Config {
  let http: HttpClient;
  let router: Router;
  const configMap: Map<string, string> = new Map<string, string>();
  let triedLogin = false;
  export let login: boolean;
  export let captcha: string;
  export const baseUrl = environment.production
    ? '/'
    : 'http://localhost:1813/';
  export const apiUrl = baseUrl + 'api/';

  export function get(key: string): string {
    return configMap.get(key);
  }

  export function getCaptcha() {
    API('newcaptcha', {}).subscribe(values => captcha = values['captcha']);
  }

  export function load(keys: string[]) {
    API('conf', { confkeys: keys }).subscribe(values =>
      Object.entries(values).forEach(([key, value]) => configMap.set(key, value as string))
    );
  }

  export function loadFirst(keys: string[], title: Title) {
    API('conf', { confkeys: keys }).subscribe(values =>
      Object.entries(values).forEach(([key, value]) => loadFirstSet(key, value as string, title))
    );
  }

  function loadFirstSet(key: string, value: string, title: Title) {
    configMap.set(key, value);
    if (key === 'title') {
      title.setTitle(value);
    }
  }

  export function API(url: string, body: any): Observable<any> {
    return http.post<any>(apiUrl + url, body);
  }

  export function setLogin(redirect: boolean) {
    if (!triedLogin) {
      Config.API('login', { username: localStorage.getItem('username'), password: localStorage.getItem('password') })
        .subscribe(values => validateLogin(values, redirect));
    } else {
      if (redirect && !login) {
        router.navigate(['/']);
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

  export function setRouter(newRouter: Router) {
    router = newRouter;
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
    const date = new Date(<string>registered);
    return ((date.getDate() <= 9 ? '0' + date.getDate() : date.getDate()) + '.' + (date.getMonth() + 1) + '.' + date.getFullYear());
  }

  export function createdDate(created: Object): string {
    const date = new Date(<number>created * 1000);
    return ((date.getDate() <= 9 ? '0' + date.getDate() : date.getDate()) + '.' + (date.getMonth() + 1) + '.' + date.getFullYear());
  }

  export function setHttp(httpClient: HttpClient) {
    http = httpClient;
  }

  export function hash(text: string): string {
    return sha3_512('gorum_' + sha3_512(text));
  }
}
