export namespace dialog {
	
	export enum DialogType {
	    info = "info",
	    warning = "warning",
	    error = "error",
	    question = "question",
	}
	export enum DialogActionType {
	    Ok = "Ok",
	    OkCancel = "OkCancel",
	    YesNo = "YesNo",
	}
	export enum DialogButton {
	    Ok = "Ok",
	    Cancel = "Cancel",
	    Yes = "Yes",
	    No = "No",
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
	    status: StatusItem;
	
	    static createFrom(source: any = {}) {
	        return new CreateCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	    }
	}

}

