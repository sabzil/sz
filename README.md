sz
==
sz(SnapZil)은 문서작업을 하다 이미지를 자르는게 귀찮아서 시작하게된 프로젝트 입니다. 그렇지만 편하지는 않습니다.
그냥 이미지 편집도구로 작업하길 권장합니다.

## 설치
`go get` 으로 설치합니다.
```
go get github.com/mitchellh/cli
go get github.com/disintegration/imaging
go get github.com/oliamb/cutter
go get github.com/sabzil/sz
```

## 사용법
```
자르기: sz crop -source="./src.jpg" -target="./a.jpg" -x=0 -y=0 -width=30 -height=30
회전: sz rotate -source="./src.jpg" -direction="left"
확대/축소: sz resize -source="./src.jpg" -width=80 -height=50 -filter="NearestNeighbor"
```

