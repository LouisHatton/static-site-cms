export interface SiteStructure {
	root: string;
	routes: SiteRoute[];
}

export interface SiteRoute {
	path: string;
	files: string[] | null;
}

export const DefaultSiteStructure: SiteStructure = {
	root: '',
	routes: []
};
