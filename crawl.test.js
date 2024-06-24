import { test, expect } from "@jest/globals";
import { normalizeURL } from "./crawl.js";


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