export function DeepCopyObj(obj: any): any {
	return JSON.parse(JSON.stringify(obj));
}
