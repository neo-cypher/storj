// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

import { Component } from 'vue';

import { BrowserObject } from '@/store/modules/objectBrowserStore';

import RedditIcon from '@poc/components/icons/share/IconReddit.vue';
import FacebookIcon from '@poc/components/icons/share/IconFacebook.vue';
import TwitterIcon from '@poc/components/icons/share/IconTwitter.vue';
import HackerNewsIcon from '@poc/components/icons/share/IconHackerNews.vue';
import LinkedInIcon from '@poc/components/icons/share/IconLinkedIn.vue';
import TelegramIcon from '@poc/components/icons/share/IconTelegram.vue';
import WhatsAppIcon from '@poc/components/icons/share/IconWhatsApp.vue';
import EmailIcon from '@poc/components/icons/share/IconEmail.vue';

import imageIcon from '@poc/assets/icon-image-tonal.svg';
import videoIcon from '@poc/assets/icon-video-tonal.svg';
import audioIcon from '@poc/assets/icon-audio-tonal.svg';
import textIcon from '@poc/assets/icon-text-tonal.svg';
import pdfIcon from '@poc/assets/icon-pdf-tonal.svg';
import zipIcon from '@poc/assets/icon-zip-tonal.svg';
import spreadsheetIcon from '@poc/assets/icon-spreadsheet-tonal.svg';
import folderIcon from '@poc/assets/icon-folder-tonal.svg';
import fileIcon from '@poc/assets/icon-file-tonal.svg';

export enum ShareOptions {
    Reddit = 'Reddit',
    Facebook = 'Facebook',
    Twitter = 'Twitter',
    HackerNews = 'Hacker News',
    LinkedIn = 'LinkedIn',
    Telegram = 'Telegram',
    WhatsApp = 'WhatsApp',
    Email = 'E-Mail',
}

export interface ShareButtonConfig {
    color: string;
    getLink: (linksharingURL: string) => string;
    icon: Component;
}

export const SHARE_BUTTON_CONFIGS: Record<ShareOptions, ShareButtonConfig> = {
    [ShareOptions.Reddit]: {
        color: '#5f99cf',
        getLink: url => `https://reddit.com/submit/?url=${url}&resubmit=true&title=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage`,
        icon: RedditIcon,
    },
    [ShareOptions.Facebook]: {
        color: '#3b5998',
        getLink: url => `https://facebook.com/sharer/sharer.php?u=${url}`,
        icon: FacebookIcon,
    },
    [ShareOptions.Twitter]: {
        color: '#55acee',
        getLink: url => `https://twitter.com/intent/tweet/?text=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage&url=${url}`,
        icon: TwitterIcon,
    },
    [ShareOptions.HackerNews]: {
        color: '#f60',
        getLink: url => `https://news.ycombinator.com/submitlink?u=${url}&t=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage`,
        icon: HackerNewsIcon,
    },
    [ShareOptions.LinkedIn]: {
        color: '#0077b5',
        getLink: url => `https://www.linkedin.com/shareArticle?mini=true&url=${url}&title=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage&summary=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage&source=${url}`,
        icon: LinkedInIcon,
    },
    [ShareOptions.Telegram]: {
        color: '#54a9eb',
        getLink: url => `https://telegram.me/share/url?text=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage&url=${url}`,
        icon: TelegramIcon,
    },
    [ShareOptions.WhatsApp]: {
        color: '#25d366',
        getLink: url => `whatsapp://send?text=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage%20${url}`,
        icon: WhatsAppIcon,
    },
    [ShareOptions.Email]: {
        color: '#777',
        getLink: url => `mailto:?subject=Shared%20using%20Storj%20Decentralized%20Cloud%20Storage&body=${url}`,
        icon: EmailIcon,
    },
};

export enum ShareType {
    File = 'File',
    Folder = 'Folder',
    Bucket = 'Bucket',
}

export enum PreviewType {
    None,
    Text,
    CSV,
    Image,
    Video,
    Audio,
    PDF,
}

export const EXTENSION_PREVIEW_TYPES = new Map<string[], PreviewType>([
    [['txt', 'md', 'json', 'xml'], PreviewType.Text],
    [['csv'], PreviewType.CSV],
    [['bmp', 'svg', 'jpg', 'jpeg', 'png', 'ico', 'gif', 'webp'], PreviewType.Image],
    [['m4v', 'mp4', 'webm', 'mov', 'mkv', 'ogv'], PreviewType.Video],
    [['m4a', 'mp3', 'wav', 'ogg', 'aac', 'flac'], PreviewType.Audio],
    [['pdf'], PreviewType.PDF],
]);

export type BrowserObjectTypeInfo = {
    title: string;
    icon: string;
};

/**
 * Contains extra information to aid in the display, filtering, and sorting of browser objects.
 */
export type BrowserObjectWrapper = {
    browserObject: BrowserObject;
    typeInfo: BrowserObjectTypeInfo;
    lowerName: string;
    ext: string;
};

export const EXTENSION_INFOS: Map<string[], BrowserObjectTypeInfo> = new Map([
    [['bmp', 'svg', 'jpg', 'jpeg', 'png', 'ico', 'gif', 'webp'], { title: 'Image', icon: imageIcon }],
    [['m4v', 'mp4', 'webm', 'mov', 'mkv', 'ogv', 'avi'], { title: 'Video', icon: videoIcon }],
    [['m4a', 'mp3', 'wav', 'ogg', 'aac', 'flac', 'aiff'], { title: 'Audio', icon: audioIcon }],
    [['txt', 'docx', 'doc', 'pages'], { title: 'Text', icon: textIcon }],
    [['pdf'], { title: 'PDF', icon: pdfIcon }],
    [['zip'], { title: 'ZIP', icon: zipIcon }],
    [['xls', 'numbers', 'csv', 'xlsx', 'tsv'], { title: 'Spreadsheet', icon: spreadsheetIcon }],
]);
export const FOLDER_INFO: BrowserObjectTypeInfo = { title: 'Folder', icon: folderIcon };
export const FILE_INFO: BrowserObjectTypeInfo = { title: 'File', icon: fileIcon };
