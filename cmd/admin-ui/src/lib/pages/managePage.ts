export interface SiteStructure {
	root: string;
	routes: SiteRoute[];
}

export interface SiteRoute {
	path: string;
	files: string[];
}

export const DefaultSiteStructure: SiteStructure = {
	root: '',
	routes: []
};
