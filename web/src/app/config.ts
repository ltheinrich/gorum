import { sha3_512 } from 'js-sha3';

export module Config {
    export function Get(parent: string, child: string): string {
        switch (parent) {
            case 'title':
                return 'Gorum';
            default:
                return null;
        }
    }

    export function API(url: string): string {
        return 'https://localhost:1813/' + url;
    }

    export function Hash(text: string): string {
        return sha3_512('secpass_' + sha3_512(text));
    }
}
