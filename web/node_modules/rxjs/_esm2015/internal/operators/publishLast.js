import { AsyncSubject } from '../AsyncSubject';
import { multicast } from './multicast';
export function publishLast() {
    return (source) => multicast(new AsyncSubject())(source);
}
