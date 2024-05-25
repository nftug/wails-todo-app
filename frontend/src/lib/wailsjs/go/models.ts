export namespace dialog {
	
	export enum DialogButton {
	    Ok = "Ok",
	    Cancel = "Cancel",
	    Yes = "Yes",
	    No = "No",
	}
	export enum DialogType {
	    Info = "info",
	    Warning = "warning",
	    Error = "error",
	    Question = "question",
	}
	export enum DialogActionType {
	    Ok = "Ok",
	    OkCancel = "OkCancel",
	    YesNo = "YesNo",
	}
	export class DialogOptions {
	    message: string;
	    title?: string;
	    type?: DialogType;
	    actionType?: DialogActionType;
	
	    static createFrom(source: any = {}) {
	        return new DialogOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	        this.title = source["title"];
	        this.type = source["type"];
	        this.actionType = source["actionType"];
	    }
	}

}

export namespace interfaces {
	
	export class CreatedResponse {
	    id: number[];
	
	    static createFrom(source: any = {}) {
	        return new CreatedResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}

}

export namespace todo {
	
	export enum StatusItem {
	    Backlog = "Backlog",
	    Todo = "Todo",
	    Doing = "Doing",
	    Done = "Done",
	}
	export class CreateCommand {
	    title: string;
	    description?: string;
	    initialStatus?: StatusItem;
	    // Go type: time
	    dueDate?: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.initialStatus = source["initialStatus"];
	        this.dueDate = this.convertValues(source["dueDate"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DetailResponse {
	    id: number[];
	    title: string;
	    description?: string;
	    status: StatusItem;
	    // Go type: time
	    statusUpdatedAt: any;
	    // Go type: time
	    dueDate?: any;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt?: any;
	
	    static createFrom(source: any = {}) {
	        return new DetailResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.statusUpdatedAt = this.convertValues(source["statusUpdatedAt"], null);
	        this.dueDate = this.convertValues(source["dueDate"], null);
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ItemResponse {
	    id: number[];
	    title: string;
	    status: StatusItem;
	    // Go type: time
	    dueDate?: any;
	
	    static createFrom(source: any = {}) {
	        return new ItemResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status = source["status"];
	        this.dueDate = this.convertValues(source["dueDate"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Query {
	    search?: string;
	    title?: string;
	    description?: string;
	    status?: StatusItem;
	    // Go type: time
	    until?: any;
	
	    static createFrom(source: any = {}) {
	        return new Query(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.search = source["search"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.until = this.convertValues(source["until"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class UpdateCommand {
	    title: string;
	    description?: string;
	    // Go type: time
	    dueDate?: any;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.dueDate = this.convertValues(source["dueDate"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class UpdateStatusCommand {
	    status: StatusItem;
	
	    static createFrom(source: any = {}) {
	        return new UpdateStatusCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	    }
	}

}
