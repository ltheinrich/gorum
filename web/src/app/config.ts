import { sha3_512 } from 'js-sha3';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export module Config {
    let http: HttpClient;
    export let login: boolean;

    export function Get(parent: string, child: string): string {
        switch (parent) {
            case 'title':
                return 'Gorum';
            default:
                return null;
        }
    }

    export function API(url: string, body: any): Observable<any> {
        return http.post<any>('https://localhost:1813/api/' + url, body);
    }

    export function setLogin() {
        Config.API('login',
            { username: localStorage.getItem('username'), password: localStorage.getItem('password') })
            .subscribe(values => login = values['valid']);
    }

    export function setHttp(httpClient: HttpClient) { http = httpClient; }

    export function Hash(text: string): string { return sha3_512('gorum_' + sha3_512(text)); }
}
