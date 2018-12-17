(require 'cl-utilities)
(require 'cl-containers)
(require 'cl-strings)


(defparameter *e-attack* 3)
(defparameter *g-attack* 3)

(defun file->2darray (path)
  (let* ((lines (read-lines path))
         (h (length lines))
         (w (length (car lines))))
    (values
     (make-array (list h w) :initial-contents (mapcar #'(lambda (l) (coerce l 'list)) lines))
     h w)))

(defun read-day15-input (path)
  (multiple-value-bind (arr h w) (file->2darray path)
    (loop for y from 0 below h do
         (loop for x from 0 below w do
              (case (aref arr y x)
                ((#\E) (setf (aref arr y x) (make-entity :type 'elf)))
                ((#\G) (setf (aref arr y x) (make-entity :type 'goblin))))))
    arr))

(defun print-cavern (c)
  (let ((h (array-dimension c 0))
        (w (array-dimension c 1)))
    (loop for y from 0 below h do
         (let ((ents '()))
           (loop for x from 0 below w do
                (let ((e (aref c y x)))
                  (if (entity-p e)
                      (progn
                        (format t "~a" (if (eq 'elf (entity-type e)) #\E #\G))
                        (push (entity-desc e) ents))
                      (format t "~a" e))))
           (format t " ~a ~%" (if ents (reverse ents) ""))))))

(defstruct entity
  (type nil)
  (hp 200 :type fixnum))

(defun entity-desc (e)
  (format nil "~a(~a)" (if (eq 'elf (entity-type e)) #\E #\G) (entity-hp e)))

(declaim (inline visit-cavern))
(defun visit-cavern (c f)
  (let ((h (array-dimension c 0))
        (w (array-dimension c 1)))
    (loop for y from 0 below h do
         (loop for x from 0 below w do
              (let ((e (aref c y x)))
                (funcall f e (make-cord :x x :y y)))))))

(defun get-all-entities-pos (c)
  (let ((pos '()))
    (flet ((collect (e p) (when (entity-p e) (push p pos))))
      (visit-cavern c #'collect))
    (reverse pos)))

(declaim (inline cavern-at)
         (optimize (speed 3) (safety 1)))
(defun cavern-at (c pos)
  (declare (optimize (speed 3) (safety 1)))
  (aref c (cord-y pos) (cord-x pos)))

(defun enemies-of (c pos)
  (let* ((all (get-all-entities-pos c))
         (type (entity-type (cavern-at c pos)))
         (enemies (remove-if #'(lambda (p) (eq type (entity-type (cavern-at c p)))) all)))
    enemies))

(defparameter *ent-neighbours* (list (make-cord :x 0  :y -1)
                                     (make-cord :x -1 :y 0)
                                     (make-cord :x 1  :y 0)
                                     (make-cord :x 0  :y 1)))

(declaim (inline select-dest))
(defun select-dest (c from)
  (enemies-of c from))

(declaim (inline cavern-free-at))
(defun cavern-free-at (c p)
  (declare (optimize (speed 3)))
  (equalp #\. (cavern-at c p)))

(defun free-neighbours (c pos)
  (remove-if-not #'(lambda (p) (cavern-free-at c p))
                 (mapcar #'(lambda (p) (move-coord pos p)) *ent-neighbours*)))

(defun find-path (c from to)
  (let ((to-visit (free-neighbours c from))
        (pred (make-hash-table :test #'equalp))
        (path '()))
    (loop for p in to-visit do
         (setf (gethash p pred) from))
    (block outer
      (loop while (not (empty-p to-visit)) do
           (let* ((next (pop to-visit))
                  (neighbours (free-neighbours c next))
                  (unseen (remove-if #'(lambda (p) (gethash p pred)) (copy-seq  neighbours))))
             (loop for p in unseen do
                  (setf (gethash p pred) next))
             (setf to-visit (append to-visit unseen))
             (when (gethash to pred)
               (return-from outer)))))
    (when (not (gethash to pred))
      (return-from find-path '()))
    (let ((current to))
      (loop while (not (equalp current from)) do
           (push current path)
           (setf current (gethash current pred))))
    path))

(defun cord-reading-score (c)
  (+ (* 10000 (cord-y c)) (cord-x c)))

(defun sort-by-reading (l)
  (sort l #'< :key #'cord-reading-score))

(defun smallest-reading (l)
  (let ((sorted (sort-by-reading l)))
    (when (consp sorted) (car sorted))))

(defun compute-paths (c start ends)
  (let ((paths (make-hash-table :test #'equalp)))
    (loop for end in ends do
         (setf (gethash (cons start end) paths) (find-path c start end)))
    paths))

(defun next-move (c pos)
  (declare (optimize (speed 3) (safety 1)))
  (let* ((enemies (enemies-of c pos))
         (in-range (remove-duplicates (mapcan #'(lambda (e) (free-neighbours c e))
                                              enemies) :test #'equalp))
         (computed-paths (compute-paths c pos in-range))
         (reachable (remove-if-not
                     #'(lambda (e) (gethash (cons pos e) computed-paths)) in-range))
         (nearest-dist (loop for e in reachable minimizing
                            (length (gethash (cons pos e) computed-paths))))
         (nearest (loop for e in reachable when (= nearest-dist (length (gethash (cons pos e) computed-paths))) collecting e))
         (nearest-step nil)
         (shortest-path most-positive-fixnum)
         (nearest-step-score most-positive-fixnum))
    (loop for n in nearest do
         (let* ((path (gethash (cons pos n) computed-paths))
                (path-len (length path))
                (first (car path))
                (score (the fixnum (cord-reading-score first))))
           (if (< path-len shortest-path)
               (progn
                 (setf shortest-path path-len
                       nearest-step first
                       nearest-step-score score))
               (progn
                 (when (and (= path-len shortest-path) (< score nearest-step-score))
                   (setf nearest-step first
                         nearest-step-score score))))))
    nearest-step))

(defun attack-power (c pos)
  (let ((type (entity-type (cavern-at c pos))))
    (if (eq type 'elf) *g-attack* *e-attack*)))

(defun attack (c p)
  (assert (entity-p (cavern-at c p)))
  (let ((copy (copy-structure (aref c (cord-y p) (cord-x p)))))
    (decf (entity-hp copy) (attack-power c p))
    (if (>= 0 (entity-hp copy))
        (progn
          (setf (aref c (cord-y p) (cord-x p)) #\.)
          (return-from attack t))
        (progn
          (setf (aref c (cord-y p) (cord-x p)) copy)
          (return-from attack nil)))))

(defun sum-hp (c type)
  (let* ((all (mapcar #'(lambda (p) (cavern-at c p)) (get-all-entities-pos c)))
         (matching (remove-if-not #'(lambda (e) (equalp type (entity-type e))) all)))
    (loop for e in matching summing (entity-hp e))))

(defun game-ended (c)
  (let ((ghp (sum-hp c 'goblin))
        (ehp (sum-hp c 'elf)))
    (when (or (= 0 ghp) (= 0 ehp))
      (max ghp ehp))))

(defun cavern-turn (c)
  (declare (optimize (speed 3) (safety 1)))
  (let* ((all (get-all-entities-pos c))
         (killed (make-hash-table :test #'equalp)))
    (loop for e in all do
         (when (and (entity-p (cavern-at c e)) (not (gethash e killed)))
           (let ((attack (pos-to-attack c e)))
             (if attack
                 (progn
                   (let ((k (attack c attack)))
                     (when k
                       (setf (gethash attack killed) t))))
                 (let ((next-pos (next-move c e)))
                   (if next-pos
                       (progn
                         (setf (aref c (cord-y next-pos) (cord-x next-pos))
                               (aref c (cord-y e) (cord-x e))
                               (aref c (cord-y e) (cord-x e))
                               #\.)
                         (let ((attack (pos-to-attack c next-pos)))
                           (when attack
                             (attack c attack))))
                       (when (game-ended c)
                         (return-from cavern-turn (values c (game-ended c) t))))))))))
  (let ((ghp (sum-hp c 'goblin))
        (ehp (sum-hp c 'elf)))
    (if (or (= 0 ghp) (= 0 ehp))
        (values c (max ghp ehp) nil)
        (values c nil nil))))

(defun count-elfs (c)
  (length (remove-if-not #'(lambda (e) (eq 'elf (entity-type e)))
                         (mapcar #'(lambda (p) (cavern-at c p)) (get-all-entities-pos c)))))

(defun cavern-turns (c turns)
  (let ((current (cl-utilities:copy-array c)))
    (loop for x from 1 to turns do
         (multiple-value-bind (n h not-full) (cavern-turn current)
           (setf current n)
           (when h
             (return-from cavern-turns (values (* (if not-full (1- x) x)  h) (count-elfs current))))))
    current))

(defun are-enemies (c a b)
  (let* ((ae (cavern-at c a))
         (be (cavern-at c b)))
    (and (entity-p ae) (entity-p be) (not (eq (entity-type ae) (entity-type be))))))

(defun adjecent-enemies (c pos)
  (remove-if-not #'(lambda (p) (are-enemies c pos p))
                 (mapcar #'(lambda (p) (move-coord pos p)) *ent-neighbours*)))

(defun pos-to-attack (c pos)
  (declare (optimize (speed 3) (safety 1)))
  (let* ((enemies (adjecent-enemies c pos))
         (lowest-hp (the fixnum (loop for e in enemies minimizing (entity-hp (cavern-at c e)))))
         (candidates (loop for e in enemies when (= lowest-hp (the fixnum (entity-hp (cavern-at c e)))) collecting e)))
    (smallest-reading candidates)))


(defun part2 (input &optional (start 4))
  (let ((elfs-start (count-elfs input)))
    (loop for attack from start do
         (let* ((*e-attack* attack))
           (multiple-value-bind (score elfs-after) (cavern-turns input 10000)
             (when (= elfs-start elfs-after)
               (return-from part2 (values score *e-attack*))))))))
