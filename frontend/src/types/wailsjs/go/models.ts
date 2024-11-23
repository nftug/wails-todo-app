export namespace enums {
	
	export enum StatusValue {
	    Backlog = "Backlog",
	    Todo = "Todo",
	    Doing = "Doing",
	    Done = "Done",
	}
	export enum TodoEvent {
	    NotifyTodo = "NotifyTodo",
	}
	export enum ErrorCode {
	    InvalidArg = "InvalidArg",
	    NotFound = "NotFound",
	}

}

export namespace interfaces {
	
	export interface CreatedResponse {
	    id: string;
	}

}

export namespace todo {
	
	export interface CreateCommand {
	    title: string;
	    description?: string;
	    initialStatus?: enums.StatusValue;
	    // Go type: time
	    dueDate?: any;
	}
	export interface DetailsResponse {
	    id: string;
	    title: string;
	    description?: string;
	    status: enums.StatusValue;
	    // Go type: time
	    statusUpdatedAt: any;
	    // Go type: time
	    dueDate?: any;
	    // Go type: time
	    notifiedAt?: any;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt?: any;
	}
	export interface ItemResponse {
	    id: string;
	    title: string;
	    description?: string;
	    status: enums.StatusValue;
	    // Go type: time
	    notifiedAt?: any;
	    // Go type: time
	    dueDate?: any;
	}
	export interface Query {
	    search?: string;
	    title?: string;
	    description?: string;
	    status?: enums.StatusValue;
	    // Go type: time
	    until?: any;
	    // Go type: time
	    after?: any;
	    isNotified?: boolean;
	}
	export interface UpdateCommand {
	    title: string;
	    description?: string;
	    // Go type: time
	    dueDate?: any;
	}
	export interface UpdateStatusCommand {
	    status: enums.StatusValue;
	}

}

