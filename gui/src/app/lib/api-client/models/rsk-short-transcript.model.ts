/* tslint:disable */
import {
  RskSynopsis,
} from '.';

export interface RskShortTranscript {
  actors?: string[];
  audioUri?: string;
  episode?: number;
  id?: string;
  incomplete?: boolean;
  name?: string;
  offsetAccuracyPcnt?: number;
  publication?: string;
  releaseDate?: string;
  series?: number;
  shortId?: string;
  summary?: string;
  synopsis?: RskSynopsis[];
  transcriptAvailable?: boolean;
  triviaAvailable?: boolean;
}
