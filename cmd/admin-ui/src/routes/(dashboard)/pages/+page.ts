import { DefaultSiteStructure, type SiteStructure } from '$lib/pages/managePage';
import { DeepCopyObj } from '$lib/utils';
import type { PageLoad } from './$types';

export const load: PageLoad = async function (request) {
	let res = await request.fetch('/api/site/structure');
	let data: SiteStructure;
	if (res.ok) {
		data = await res.json();
	} else {
		data = DeepCopyObj(DefaultSiteStructure);
	}

	return {
		siteStructure: data
	};
};
