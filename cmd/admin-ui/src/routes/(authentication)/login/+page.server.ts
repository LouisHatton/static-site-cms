// import { } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { CookieSerializeOptions } from 'cookie';

export const load: PageServerLoad = async function (request) {
	// console.log(request.cookies.getAll());
	let cookieOptions: CookieSerializeOptions = {
		httpOnly: true
	};
	request.cookies.set('test', 'this is a test cookie from the svelte server', cookieOptions);
};
