export class Storage {
  private value: any;

  getValue(): any {
    return this.value;
  }

  setValue(mediator: Mediator, storageName: string, value: any) {
    this.value = value;
    mediator.notifyObservers(storageName);
  }
}

export class Mediator {
  private storageMap: {
    [key: string]: Storage;
  } = {};

  private observers: {
    [key: string]: () => void;
  } = {};

  setValue(storageName: string, value: any) {
    const storage = new Storage();
    this.storageMap[storageName] = storage;
    storage.setValue(this, storageName, value);
  }

  getValue(storageName: string) {
    return this.storageMap[storageName].getValue();
  }
  addObserver(storageName: string, observer: () => void): void {
    this.observers[storageName] = observer;
  }

  notifyObservers(eventName: string): void {
    const observer = this.observers[eventName];
    observer();
  }
}
