import { Config } from './config';

export module Language {
    const languageMap: Map<string, string> = new Map<string, string>();

    export function get(name: string): string {
        return languageMap.get(name);
    }

    export function loadLanguage(language: string) {
        Config.API('lang', {}).subscribe(values =>
            Object.entries(values[language]).forEach(([key, value]) => languageMap.set(key, value as string)));
    }
}
