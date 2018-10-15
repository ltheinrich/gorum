import { sha3_512 } from 'js-sha3';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

export module Config {
    let http: HttpClient;
    const configMap: Map<string, string> = new Map<string, string>();
    export let login: boolean;

    export function get(key: string): string {
        return configMap.get(key);
    }

    export function load(key: string) {
        API('conf', { confkey: key }).subscribe(values => configMap.set(key, values[key]));
    }

    export function API(url: string, body: any): Observable<any> {
        const apiUrl = environment.production ? '/api/' : 'https://localhost:1813/api/';
        return http.post<any>(apiUrl + url, body);
    }

    export function setLogin() {
        Config.API('login',
            { username: localStorage.getItem('username'), password: localStorage.getItem('password') })
            .subscribe(values => login = values['valid']);
    }

    export function setHttp(httpClient: HttpClient) { http = httpClient; }

    export function hash(text: string): string { return sha3_512('gorum_' + sha3_512(text)); }
}
