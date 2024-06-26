import { test, expect } from "@jest/globals";
import { normalizeURL, getURLsFromHTML } from "./crawl.js";

describe('normalizeURL', () => {
    it('should normalize a simple URL', () => {
        const input = 'http://example.com';
        const expected = 'example.com';
        expect(normalizeURL(input)).toBe(expected);
    });

    it('should remove trailing slash', () => {
        const input = 'http://example.com/';
        const expected = 'example.com';
        expect(normalizeURL(input)).toBe(expected);
    });

    it('should handle URLs with paths', () => {
        const input = 'http://example.com/path/';
        const expected = 'example.com/path';
        expect(normalizeURL(input)).toBe(expected);
    });

    it('should handle URLs with query parameters', () => {
        const input = 'http://example.com/path/?query=1';
        const expected = 'example.com/path/?query=1';
        expect(normalizeURL(input)).toBe(expected);
    });

    it('should handle URLs with hashes', () => {
        const input = 'http://example.com/path/#hash';
        const expected = 'example.com/path/#hash';
        expect(normalizeURL(input)).toBe(expected);
    });

    it('should throw an error for invalid URLs', () => {
        const input = 'invalid-url';
        expect(() => normalizeURL(input)).toThrow('Invalid URL');
    });

    it('should handle HTTPS URLs', () => {
        const input = 'https://example.com';
        const expected = 'example.com';
        expect(normalizeURL(input)).toBe(expected);
    });
});

describe('getURLsFromHTML', () => {
    it('should return an empty array for an empty HTML body', () => {
        const input = '';
        const baseURL = 'http://example.com';
        const expected = [];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an empty array for an HTML body without links', () => {
        const input = '<html><body></body></html>';
        const baseURL = 'http://example.com';
        const expected = [];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body', () => {
        const input = '<html><body><a href="http://example.com">Link</a></body></html>';
        const baseURL = 'http://example.com';
        const expected = ['http://example.com/'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with multiple links', () => {
        const input = '<html><body><a href="http://example.com">Link 1</a><a href="http://example.com/path">Link 2</a></body></html>';
        const baseURL = 'http://example.com';
        const expected = ['http://example.com/', 'http://example.com/path'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with relative links', () => {
        const input = '<html><body><a href="/path">Link</a></body></html>';
        const baseURL = 'http://example.com';
        const expected = ['http://example.com/path'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with links to other domains', () => {
        const input = '<html><body><a href="http://example.com">Link</a><a href="http://example.com/path">Link</a></body></html>';
        const baseURL = 'http://other.com';
        const expected = ['http://example.com/', 'http://example.com/path'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with links to other domains and relative links', () => {
        const input = '<html><body><a href="http://example.com">Link</a><a href="/path">Link</a></body></html>';
        const baseURL = 'http://other.com';
        const expected = ['http://example.com/', 'http://other.com/path'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with links to other domains and relative links with query parameters', () => {
        const input = '<html><body><a href="http://example.com">Link</a><a href="/path?query=1">Link</a></body></html>';
        const baseURL = 'http://other.com';
        const expected = ['http://example.com/', 'http://other.com/path?query=1'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });

    it('should return an array of URLs from an HTML body with links to other domains and relative links with hashes', () => {
        const input = '<html><body><a href="http://example.com">Link</a><a href="/path#hash">Link</a></body></html>';
        const baseURL = 'http://other.com';
        const expected = ['http://example.com/', 'http://other.com/path#hash'];
        expect(getURLsFromHTML(input, baseURL)).toEqual(expected);
    });
});
