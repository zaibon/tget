#!/usr/bin/python
import json
from os.path import dirname, join

def main():
    path = join(dirname(__file__), 't411_terms.json')
    with open(path) as f:
        blob = json.load(f)

    out = {
        'episodes': [],
        'seasons': [],
        'languages': [],
    }

    for num, episode in blob['637']['46']['terms'].items():
        ep_nr = episode.lstrip('Episode ')
        try:
            obj = {'key': int(ep_nr), 'value': int(num)}
            out['episodes'].append(obj)
        except ValueError:
            continue

    for num, season in blob['637']['45']['terms'].items():
        season_nr = season.lstrip('Saison ')
        try:
            obj = {'key': int(season_nr), 'value': int(num)}
            out['seasons'].append(obj)
        except ValueError:
            continue

    for num, language in blob['637']['17']['terms'].items():
        try:
            obj = {'key': language, 'value': int(num)}
            out['languages'].append(obj)
        except ValueError:
            continue

    with open('mapping.json', 'w+') as f:
        json.dump(out, f, indent=4)

if __name__ == '__main__':
    main()
