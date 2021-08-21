export class Idol {
    private name: string;
    private cover: string;
    private description: string;
    private id: string;

    constructor(name: string, description: string, cover?: string, id?: string){
        this.name = name;
        this.description = description;
        this.cover = cover ? cover : "http://via.placeholder.com/185x287";
        this.id = id ? id : "999";
    }

    public getName(){
        return this.name;
    }

    public getCover(){
        return this.cover;
    }

    public getDescription(){
        return this.description;
    }

    public setName(name: string){
        this.name = name;
    }

    public setCover(cover: string){
        this.cover = cover;
    }

    public setDescription(description: string){
        this.description = description;
    }

    public addId(id: string){
        this.id = id;
    }
}