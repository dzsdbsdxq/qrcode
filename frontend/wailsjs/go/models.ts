export namespace main {
	
	export class QrLogoData {
	    filePath: string;
	    fileName: string;
	
	    static createFrom(source: any = {}) {
	        return new QrLogoData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filePath = source["filePath"];
	        this.fileName = source["fileName"];
	    }
	}

}

